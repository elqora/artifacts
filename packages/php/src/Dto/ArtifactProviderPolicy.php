<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactProviderPolicy
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('actors', $data)) { throw new InvalidArgumentException('Missing required field actors for ArtifactProviderPolicy.'); }
        return new self($data);
    }

    /** @return array<mixed> */
    public function actors(): array
    {
        return $this->data['actors'];
    }

    public function hasActors(): bool
    {
        return array_key_exists('actors', $this->data);
    }

    public function mode(): ?string
    {
        return $this->data['mode'] ?? null;
    }

    public function hasMode(): bool
    {
        return array_key_exists('mode', $this->data);
    }

    public function delegation(): ?string
    {
        return $this->data['delegation'] ?? null;
    }

    public function hasDelegation(): bool
    {
        return array_key_exists('delegation', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['actors' => true, 'mode' => true, 'delegation' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
