<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class LocalArtifactSource
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('type', $data)) { throw new InvalidArgumentException('Missing required field type for LocalArtifactSource.'); }
        if (!array_key_exists('localId', $data)) { throw new InvalidArgumentException('Missing required field localId for LocalArtifactSource.'); }
        if (!array_key_exists('syncState', $data)) { throw new InvalidArgumentException('Missing required field syncState for LocalArtifactSource.'); }
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

    public function localId(): string
    {
        return $this->data['localId'];
    }

    public function hasLocalId(): bool
    {
        return array_key_exists('localId', $this->data);
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

    public function syncState(): string
    {
        return $this->data['syncState'];
    }

    public function hasSyncState(): bool
    {
        return array_key_exists('syncState', $this->data);
    }

    public function remoteVersionId(): ?string
    {
        return $this->data['remoteVersionId'] ?? null;
    }

    public function hasRemoteVersionId(): bool
    {
        return array_key_exists('remoteVersionId', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['type' => true, 'localId' => true, 'filename' => true, 'mediaType' => true, 'size' => true, 'syncState' => true, 'remoteVersionId' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
