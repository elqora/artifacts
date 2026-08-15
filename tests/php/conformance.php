<?php

declare(strict_types=1);

require __DIR__ . '/../../packages/php/src/Vocabulary.php';
require __DIR__ . '/../../packages/php/src/Protocol.php';

use Elqora\Artifact\Protocol;
use Elqora\Artifact\Vocabulary;

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
