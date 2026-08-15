<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class AudioArtifactValueSchema
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('valueType', $data)) { throw new InvalidArgumentException('Missing required field valueType for AudioArtifactValueSchema.'); }
        return new self($data);
    }

    public function minFiles(): ?int
    {
        return $this->data['minFiles'] ?? null;
    }

    public function hasMinFiles(): bool
    {
        return array_key_exists('minFiles', $this->data);
    }

    public function maxFiles(): ?int
    {
        return $this->data['maxFiles'] ?? null;
    }

    public function hasMaxFiles(): bool
    {
        return array_key_exists('maxFiles', $this->data);
    }

    /** @return array<mixed>|null */
    public function acceptedMimeTypes(): ?array
    {
        return $this->data['acceptedMimeTypes'] ?? null;
    }

    public function hasAcceptedMimeTypes(): bool
    {
        return array_key_exists('acceptedMimeTypes', $this->data);
    }

    public function maxSizeBytes(): ?int
    {
        return $this->data['maxSizeBytes'] ?? null;
    }

    public function hasMaxSizeBytes(): bool
    {
        return array_key_exists('maxSizeBytes', $this->data);
    }

    public function valueType(): string
    {
        return $this->data['valueType'];
    }

    public function hasValueType(): bool
    {
        return array_key_exists('valueType', $this->data);
    }

    public function maxDurationSeconds(): float|int|null
    {
        return $this->data['maxDurationSeconds'] ?? null;
    }

    public function hasMaxDurationSeconds(): bool
    {
        return array_key_exists('maxDurationSeconds', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['minFiles' => true, 'maxFiles' => true, 'acceptedMimeTypes' => true, 'maxSizeBytes' => true, 'valueType' => true, 'maxDurationSeconds' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
