<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactRetentionPolicyDuration
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('policy', $data)) { throw new InvalidArgumentException('Missing required field policy for ArtifactRetentionPolicyDuration.'); }
        if (!array_key_exists('days', $data)) { throw new InvalidArgumentException('Missing required field days for ArtifactRetentionPolicyDuration.'); }
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

    public function days(): int
    {
        return $this->data['days'];
    }

    public function hasDays(): bool
    {
        return array_key_exists('days', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['policy' => true, 'days' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
