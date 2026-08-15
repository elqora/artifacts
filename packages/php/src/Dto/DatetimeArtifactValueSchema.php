<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class DatetimeArtifactValueSchema
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('valueType', $data)) { throw new InvalidArgumentException('Missing required field valueType for DatetimeArtifactValueSchema.'); }
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

    public function minimum(): ?string
    {
        return $this->data['minimum'] ?? null;
    }

    public function hasMinimum(): bool
    {
        return array_key_exists('minimum', $this->data);
    }

    public function maximum(): ?string
    {
        return $this->data['maximum'] ?? null;
    }

    public function hasMaximum(): bool
    {
        return array_key_exists('maximum', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['valueType' => true, 'minimum' => true, 'maximum' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
