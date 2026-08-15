<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactEncryptionPolicy
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('required', $data)) { throw new InvalidArgumentException('Missing required field required for ArtifactEncryptionPolicy.'); }
        return new self($data);
    }

    public function required(): bool
    {
        return $this->data['required'];
    }

    public function hasRequired(): bool
    {
        return array_key_exists('required', $this->data);
    }

    public function level(): ?string
    {
        return $this->data['level'] ?? null;
    }

    public function hasLevel(): bool
    {
        return array_key_exists('level', $this->data);
    }

    public function keyScope(): ?string
    {
        return $this->data['keyScope'] ?? null;
    }

    public function hasKeyScope(): bool
    {
        return array_key_exists('keyScope', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['required' => true, 'level' => true, 'keyScope' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
