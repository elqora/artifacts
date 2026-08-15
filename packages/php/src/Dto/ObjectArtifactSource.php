<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ObjectArtifactSource
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('type', $data)) { throw new InvalidArgumentException('Missing required field type for ObjectArtifactSource.'); }
        if (!array_key_exists('objectId', $data)) { throw new InvalidArgumentException('Missing required field objectId for ObjectArtifactSource.'); }
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

    public function objectId(): string
    {
        return $this->data['objectId'];
    }

    public function hasObjectId(): bool
    {
        return array_key_exists('objectId', $this->data);
    }

    public function filename(): ?string
    {
        return $this->data['filename'] ?? null;
    }

    public function hasFilename(): bool
    {
        return array_key_exists('filename', $this->data);
    }

    public function mediaType(): ?string
    {
        return $this->data['mediaType'] ?? null;
    }

    public function hasMediaType(): bool
    {
        return array_key_exists('mediaType', $this->data);
    }

    public function size(): ?int
    {
        return $this->data['size'] ?? null;
    }

    public function hasSize(): bool
    {
        return array_key_exists('size', $this->data);
    }

    public function storageProvider(): ?string
    {
        return $this->data['storageProvider'] ?? null;
    }

    public function hasStorageProvider(): bool
    {
        return array_key_exists('storageProvider', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['type' => true, 'objectId' => true, 'filename' => true, 'mediaType' => true, 'size' => true, 'storageProvider' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
