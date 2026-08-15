<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactValidationPolicy
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        // This schema has no required fields.
        return new self($data);
    }

    public function mode(): ?string
    {
        return $this->data['mode'] ?? null;
    }

    public function hasMode(): bool
    {
        return array_key_exists('mode', $this->data);
    }

    /** @return array<mixed>|null */
    public function rules(): ?array
    {
        return $this->data['rules'] ?? null;
    }

    public function hasRules(): bool
    {
        return array_key_exists('rules', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['mode' => true, 'rules' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
