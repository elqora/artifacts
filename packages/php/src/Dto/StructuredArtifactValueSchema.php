<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class StructuredArtifactValueSchema
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('valueType', $data)) { throw new InvalidArgumentException('Missing required field valueType for StructuredArtifactValueSchema.'); }
        return new self($data);
    }

    public function valueType(): string
    {
        return $this->data['valueType'];
    }

    public function hasValueType(): bool
    {
        return array_key_exists('valueType', $this->data);
    }

    /** @return array<mixed>|null */
    public function jsonSchema(): ?array
    {
        return $this->data['jsonSchema'] ?? null;
    }

    public function hasJsonSchema(): bool
    {
        return array_key_exists('jsonSchema', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['valueType' => true, 'jsonSchema' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
