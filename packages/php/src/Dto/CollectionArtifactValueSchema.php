<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class CollectionArtifactValueSchema
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('valueType', $data)) { throw new InvalidArgumentException('Missing required field valueType for CollectionArtifactValueSchema.'); }
        if (!array_key_exists('itemSchema', $data)) { throw new InvalidArgumentException('Missing required field itemSchema for CollectionArtifactValueSchema.'); }
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

    /** @return array<mixed> */
    public function itemSchema(): array
    {
        return $this->data['itemSchema'];
    }

    public function hasItemSchema(): bool
    {
        return array_key_exists('itemSchema', $this->data);
    }

    public function minItems(): ?int
    {
        return $this->data['minItems'] ?? null;
    }

    public function hasMinItems(): bool
    {
        return array_key_exists('minItems', $this->data);
    }

    public function maxItems(): ?int
    {
        return $this->data['maxItems'] ?? null;
    }

    public function hasMaxItems(): bool
    {
        return array_key_exists('maxItems', $this->data);
    }

    public function uniqueItems(): ?bool
    {
        return $this->data['uniqueItems'] ?? null;
    }

    public function hasUniqueItems(): bool
    {
        return array_key_exists('uniqueItems', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['valueType' => true, 'itemSchema' => true, 'minItems' => true, 'maxItems' => true, 'uniqueItems' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
