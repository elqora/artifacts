<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactIntegrity
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('algorithm', $data)) { throw new InvalidArgumentException('Missing required field algorithm for ArtifactIntegrity.'); }
        if (!array_key_exists('hash', $data)) { throw new InvalidArgumentException('Missing required field hash for ArtifactIntegrity.'); }
        return new self($data);
    }

    public function algorithm(): string
    {
        return $this->data['algorithm'];
    }

    public function hasAlgorithm(): bool
    {
        return array_key_exists('algorithm', $this->data);
    }

    public function hash(): string
    {
        return $this->data['hash'];
    }

    public function hasHash(): bool
    {
        return array_key_exists('hash', $this->data);
    }

    public function size(): ?int
    {
        return $this->data['size'] ?? null;
    }

    public function hasSize(): bool
    {
        return array_key_exists('size', $this->data);
    }

    public function verifiedAt(): ?string
    {
        return $this->data['verifiedAt'] ?? null;
    }

    public function hasVerifiedAt(): bool
    {
        return array_key_exists('verifiedAt', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['algorithm' => true, 'hash' => true, 'size' => true, 'verifiedAt' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
