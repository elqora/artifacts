<?php

declare(strict_types=1);

namespace Elqora\Artifact;

use JsonException;

/**
 * Lossless wire helpers. Domain records are associative arrays described in
 * Contracts.php so unknown additive fields and opaque provider values survive.
 */
final class Protocol
{
    /** @return array<string, mixed> */
    public static function decode(string $json): array
    {
        $value = json_decode($json, true, 512, JSON_THROW_ON_ERROR);
        if (!is_array($value) || array_is_list($value)) {
            throw new JsonException('An Artifact Protocol record must be a JSON object.');
        }
        return $value;
    }

    /** @param array<string, mixed> $record */
    public static function encode(array $record): string
    {
        return json_encode($record, JSON_THROW_ON_ERROR | JSON_UNESCAPED_SLASHES | JSON_PRESERVE_ZERO_FRACTION);
    }

    /** @param array<string, mixed> $record */
    public static function assertProtocolVersion(array $record): void
    {
        if (($record['schemaVersion'] ?? null) !== Vocabulary::PROTOCOL_VERSION) {
            throw new JsonException('Unsupported or missing Artifact Protocol schemaVersion.');
        }
    }
}
