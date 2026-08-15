<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactRequirementPolicy
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('mode', $data)) { throw new InvalidArgumentException('Missing required field mode for ArtifactRequirementPolicy.'); }
        return new self($data);
    }

    public function mode(): string
    {
        return $this->data['mode'];
    }

    public function hasMode(): bool
    {
        return array_key_exists('mode', $this->data);
    }

    /** @return array<mixed>|null */
    public function condition(): ?array
    {
        return $this->data['condition'] ?? null;
    }

    public function hasCondition(): bool
    {
        return array_key_exists('condition', $this->data);
    }

    /** @return array<mixed>|null */
    public function blocks(): ?array
    {
        return $this->data['blocks'] ?? null;
    }

    public function hasBlocks(): bool
    {
        return array_key_exists('blocks', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['mode' => true, 'condition' => true, 'blocks' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
