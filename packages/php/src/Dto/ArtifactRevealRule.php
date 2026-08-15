<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactRevealRule
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('actors', $data)) { throw new InvalidArgumentException('Missing required field actors for ArtifactRevealRule.'); }
        if (!array_key_exists('representation', $data)) { throw new InvalidArgumentException('Missing required field representation for ArtifactRevealRule.'); }
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

    /** @return array<mixed>|null */
    public function when(): ?array
    {
        return $this->data['when'] ?? null;
    }

    public function hasWhen(): bool
    {
        return array_key_exists('when', $this->data);
    }

    public function representation(): string
    {
        return $this->data['representation'];
    }

    public function hasRepresentation(): bool
    {
        return array_key_exists('representation', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['actors' => true, 'when' => true, 'representation' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
