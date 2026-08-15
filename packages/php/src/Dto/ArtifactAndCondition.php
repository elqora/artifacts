<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactAndCondition
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('kind', $data)) { throw new InvalidArgumentException('Missing required field kind for ArtifactAndCondition.'); }
        if (!array_key_exists('conditions', $data)) { throw new InvalidArgumentException('Missing required field conditions for ArtifactAndCondition.'); }
        return new self($data);
    }

    public function kind(): string
    {
        return $this->data['kind'];
    }

    public function hasKind(): bool
    {
        return array_key_exists('kind', $this->data);
    }

    /** @return array<mixed> */
    public function conditions(): array
    {
        return $this->data['conditions'];
    }

    public function hasConditions(): bool
    {
        return array_key_exists('conditions', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['kind' => true, 'conditions' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
