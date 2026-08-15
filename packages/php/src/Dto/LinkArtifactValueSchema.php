<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class LinkArtifactValueSchema
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('valueType', $data)) { throw new InvalidArgumentException('Missing required field valueType for LinkArtifactValueSchema.'); }
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
    public function allowedSchemes(): ?array
    {
        return $this->data['allowedSchemes'] ?? null;
    }

    public function hasAllowedSchemes(): bool
    {
        return array_key_exists('allowedSchemes', $this->data);
    }

    /** @return array<mixed>|null */
    public function allowedHosts(): ?array
    {
        return $this->data['allowedHosts'] ?? null;
    }

    public function hasAllowedHosts(): bool
    {
        return array_key_exists('allowedHosts', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['valueType' => true, 'allowedSchemes' => true, 'allowedHosts' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
