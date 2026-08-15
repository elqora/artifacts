<?php

declare(strict_types=1);

require __DIR__ . '/../../packages/php/src/Vocabulary.php';
require __DIR__ . '/../../packages/php/src/Protocol.php';

spl_autoload_register(static function (string $class): void {
    $prefixes = [
        'Elqora\\Artifact\\GitHub\\' => __DIR__ . '/../../extensions/github/php/',
        'Elqora\\Artifact\\' => __DIR__ . '/../../packages/php/src/',
    ];
    foreach ($prefixes as $prefix => $base) {
        if (!str_starts_with($class, $prefix)) continue;
        $file = $base . str_replace('\\', '/', substr($class, strlen($prefix))) . '.php';
        if (is_file($file)) require $file;
        return;
    }
});

use Elqora\Artifact\Protocol;
use Elqora\Artifact\Vocabulary;
use Elqora\Artifact\Dto\Artifact;
use Elqora\Artifact\Dto\ArtifactLink;
use Elqora\Artifact\Dto\ArtifactRequirement;
use Elqora\Artifact\Dto\ArtifactSpec;
use Elqora\Artifact\Dto\ArtifactSpecSnapshot;
use Elqora\Artifact\Dto\ArtifactSubmission;
use Elqora\Artifact\Dto\ArtifactVerification;
use Elqora\Artifact\Dto\ArtifactVersion;
use Elqora\Artifact\DtoFactory;

$root = dirname(__DIR__, 2);
$fixtureDirectories = ['artifact', 'version', 'link', 'specification', 'runtime', 'condition', 'policy', 'value-schema', 'provider'];
$count = 0;
foreach ($fixtureDirectories as $directory) {
    $iterator = new RecursiveIteratorIterator(new RecursiveDirectoryIterator($root . '/tests/fixtures/' . $directory));
    foreach ($iterator as $file) {
        if (!$file->isFile() || $file->getExtension() !== 'json') continue;
        $json = file_get_contents($file->getPathname());
        if ($json === false) throw new RuntimeException('Cannot read ' . $file->getPathname());
        $record = Protocol::decode($json);
        if (array_key_exists('schemaVersion', $record)) Protocol::assertProtocolVersion($record);
        $roundTrip = Protocol::decode(Protocol::encode($record));
        if ($record !== $roundTrip) throw new RuntimeException('Semantic round trip changed ' . $file->getPathname());
        $count++;
    }
}

$protocol = json_decode(file_get_contents($root . '/spec/protocol/protocol.schema.json'), true, 512, JSON_THROW_ON_ERROR);
if (Vocabulary::ARTIFACT_VALUE_TYPES !== $protocol['$defs']['artifactValueType']['enum']) {
    throw new RuntimeException('Generated PHP value-type vocabulary is stale.');
}

fwrite(STDOUT, "PHP conformance: {$count} canonical records round-tripped without semantic change.\n");

$dtoCases = [
    [Artifact::class, 'artifact/basic.json'],
    [ArtifactVersion::class, 'version/provider-reference.json'],
    [ArtifactLink::class, 'link/pinned-version.json'],
    [ArtifactSpec::class, 'specification/delivery-evidence-v3.json'],
    [ArtifactSpecSnapshot::class, 'specification/delivery-evidence-v3-snapshot.json'],
    [ArtifactRequirement::class, 'specification/implementation-requirement.json'],
    [ArtifactSubmission::class, 'runtime/delivery-submission.json'],
    [ArtifactVerification::class, 'runtime/client-verification.json'],
];
foreach ($dtoCases as [$class, $relative]) {
    $data = json_decode(file_get_contents($root . '/tests/fixtures/' . $relative), true, 512, JSON_THROW_ON_ERROR);
    $data['futureField'] = ['preserved' => true];
    $dto = $class::fromArray($data);
    if ($dto->toArray() !== $data) throw new RuntimeException("DTO round trip changed {$relative}");
    if (!isset($dto->unknownFields()['futureField'])) throw new RuntimeException("DTO dropped unknown field for {$relative}");
}

$source = DtoFactory::source(['type' => 'provider', 'provider' => 'forge', 'reference' => ['opaque' => true]]);
if ($source->toArray()['reference'] !== ['opaque' => true]) throw new RuntimeException('Provider reference was not preserved.');
$condition = DtoFactory::condition(json_decode(file_get_contents($root . '/tests/fixtures/condition/nested.json'), true, 512, JSON_THROW_ON_ERROR));
if ($condition->toArray()['kind'] !== 'and') throw new RuntimeException('Condition factory selected the wrong DTO.');
$valueSchema = DtoFactory::valueSchema(json_decode(file_get_contents($root . '/tests/fixtures/value-schema/nested-collection.json'), true, 512, JSON_THROW_ON_ERROR));
if ($valueSchema->toArray()['valueType'] !== 'collection') throw new RuntimeException('Value schema factory selected the wrong DTO.');

$githubSource = json_decode(file_get_contents($root . '/tests/fixtures/provider/github-pull-request.json'), true, 512, JSON_THROW_ON_ERROR);
$githubReference = Elqora\Artifact\GitHub\DtoFactory::reference($githubSource['reference']);
if ($githubReference->toArray() !== $githubSource['reference']) throw new RuntimeException('GitHub DTO changed its reference.');

$submissionBase = ['schemaVersion' => '1.0', 'id' => 'sub_presence', 'artifactId' => 'art_1', 'submittedBy' => ['type' => 'system'], 'submittedAt' => '2026-08-15T09:00:00Z'];
$missingValue = ArtifactSubmission::fromArray($submissionBase);
$nullValue = ArtifactSubmission::fromArray([...$submissionBase, 'value' => null]);
if ($missingValue->hasValue() || !$nullValue->hasValue() || $nullValue->value() !== null) {
    throw new RuntimeException('PHP DTO did not preserve missing versus explicit null.');
}

fwrite(STDOUT, "PHP DTO conformance: generated DTOs, discriminated factories, and unknown-field preservation passed.\n");
