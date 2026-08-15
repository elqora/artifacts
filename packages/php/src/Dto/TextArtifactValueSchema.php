<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class TextArtifactValueSchema
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('valueType', $data)) { throw new InvalidArgumentException('Missing required field valueType for TextArtifactValueSchema.'); }
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

    public function minLength(): ?int
    {
        return $this->data['minLength'] ?? null;
    }

    public function hasMinLength(): bool
    {
        return array_key_exists('minLength', $this->data);
    }

    public function maxLength(): ?int
    {
        return $this->data['maxLength'] ?? null;
    }

    public function hasMaxLength(): bool
    {
        return array_key_exists('maxLength', $this->data);
    }

    public function multiline(): ?bool
    {
        return $this->data['multiline'] ?? null;
    }

    public function hasMultiline(): bool
    {
        return array_key_exists('multiline', $this->data);
    }

    public function pattern(): ?string
    {
        return $this->data['pattern'] ?? null;
    }

    public function hasPattern(): bool
    {
        return array_key_exists('pattern', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['valueType' => true, 'minLength' => true, 'maxLength' => true, 'multiline' => true, 'pattern' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
