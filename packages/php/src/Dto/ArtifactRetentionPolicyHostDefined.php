<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactRetentionPolicyHostDefined
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('policy', $data)) { throw new InvalidArgumentException('Missing required field policy for ArtifactRetentionPolicyHostDefined.'); }
        if (!array_key_exists('key', $data)) { throw new InvalidArgumentException('Missing required field key for ArtifactRetentionPolicyHostDefined.'); }
        return new self($data);
    }

    public function policy(): string
    {
        return $this->data['policy'];
    }

    public function hasPolicy(): bool
    {
        return array_key_exists('policy', $this->data);
    }

    public function key(): string
    {
        return $this->data['key'];
    }

    public function hasKey(): bool
    {
        return array_key_exists('key', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['policy' => true, 'key' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
