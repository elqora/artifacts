<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class CurrencyArtifactValueSchema
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('valueType', $data)) { throw new InvalidArgumentException('Missing required field valueType for CurrencyArtifactValueSchema.'); }
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
    public function currencies(): ?array
    {
        return $this->data['currencies'] ?? null;
    }

    public function hasCurrencies(): bool
    {
        return array_key_exists('currencies', $this->data);
    }

    public function minimumMinorUnits(): ?int
    {
        return $this->data['minimumMinorUnits'] ?? null;
    }

    public function hasMinimumMinorUnits(): bool
    {
        return array_key_exists('minimumMinorUnits', $this->data);
    }

    public function maximumMinorUnits(): ?int
    {
        return $this->data['maximumMinorUnits'] ?? null;
    }

    public function hasMaximumMinorUnits(): bool
    {
        return array_key_exists('maximumMinorUnits', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['valueType' => true, 'currencies' => true, 'minimumMinorUnits' => true, 'maximumMinorUnits' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
