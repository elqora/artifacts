<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ReferenceArtifactValueSchema
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('valueType', $data)) { throw new InvalidArgumentException('Missing required field valueType for ReferenceArtifactValueSchema.'); }
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
    public function providers(): ?array
    {
        return $this->data['providers'] ?? null;
    }

    public function hasProviders(): bool
    {
        return array_key_exists('providers', $this->data);
    }

    /** @return array<mixed>|null */
    public function resourceTypes(): ?array
    {
        return $this->data['resourceTypes'] ?? null;
    }

    public function hasResourceTypes(): bool
    {
        return array_key_exists('resourceTypes', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['valueType' => true, 'providers' => true, 'resourceTypes' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
