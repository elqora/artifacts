<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class InlineArtifactSource
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('type', $data)) { throw new InvalidArgumentException('Missing required field type for InlineArtifactSource.'); }
        if (!array_key_exists('value', $data)) { throw new InvalidArgumentException('Missing required field value for InlineArtifactSource.'); }
        return new self($data);
    }

    public function type(): string
    {
        return $this->data['type'];
    }

    public function hasType(): bool
    {
        return array_key_exists('type', $this->data);
    }

    public function value(): mixed
    {
        return $this->data['value'];
    }

    public function hasValue(): bool
    {
        return array_key_exists('value', $this->data);
    }

    public function mediaType(): ?string
    {
        return $this->data['mediaType'] ?? null;
    }

    public function hasMediaType(): bool
    {
        return array_key_exists('mediaType', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['type' => true, 'value' => true, 'mediaType' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
