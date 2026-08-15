<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class LocationArtifactValueSchema
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('valueType', $data)) { throw new InvalidArgumentException('Missing required field valueType for LocationArtifactValueSchema.'); }
        if (!array_key_exists('mode', $data)) { throw new InvalidArgumentException('Missing required field mode for LocationArtifactValueSchema.'); }
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

    public function mode(): string
    {
        return $this->data['mode'];
    }

    public function hasMode(): bool
    {
        return array_key_exists('mode', $this->data);
    }

    public function requireCoordinates(): ?bool
    {
        return $this->data['requireCoordinates'] ?? null;
    }

    public function hasRequireCoordinates(): bool
    {
        return array_key_exists('requireCoordinates', $this->data);
    }

    public function allowManualEntry(): ?bool
    {
        return $this->data['allowManualEntry'] ?? null;
    }

    public function hasAllowManualEntry(): bool
    {
        return array_key_exists('allowManualEntry', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['valueType' => true, 'mode' => true, 'requireCoordinates' => true, 'allowManualEntry' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
