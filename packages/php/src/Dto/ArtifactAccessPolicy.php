<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactAccessPolicy
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        // This schema has no required fields.
        return new self($data);
    }

    /** @return array<mixed>|null */
    public function read(): ?array
    {
        return $this->data['read'] ?? null;
    }

    public function hasRead(): bool
    {
        return array_key_exists('read', $this->data);
    }

    /** @return array<mixed>|null */
    public function write(): ?array
    {
        return $this->data['write'] ?? null;
    }

    public function hasWrite(): bool
    {
        return array_key_exists('write', $this->data);
    }

    /** @return array<mixed>|null */
    public function submit(): ?array
    {
        return $this->data['submit'] ?? null;
    }

    public function hasSubmit(): bool
    {
        return array_key_exists('submit', $this->data);
    }

    /** @return array<mixed>|null */
    public function verify(): ?array
    {
        return $this->data['verify'] ?? null;
    }

    public function hasVerify(): bool
    {
        return array_key_exists('verify', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['read' => true, 'write' => true, 'submit' => true, 'verify' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
