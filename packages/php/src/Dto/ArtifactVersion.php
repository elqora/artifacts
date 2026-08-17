<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactVersion
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('schemaVersion', $data)) { throw new InvalidArgumentException('Missing required field schemaVersion for ArtifactVersion.'); }
        if (!array_key_exists('id', $data)) { throw new InvalidArgumentException('Missing required field id for ArtifactVersion.'); }
        if (!array_key_exists('artifactId', $data)) { throw new InvalidArgumentException('Missing required field artifactId for ArtifactVersion.'); }
        if (!array_key_exists('version', $data)) { throw new InvalidArgumentException('Missing required field version for ArtifactVersion.'); }
        if (!array_key_exists('source', $data)) { throw new InvalidArgumentException('Missing required field source for ArtifactVersion.'); }
        if (!array_key_exists('createdBy', $data)) { throw new InvalidArgumentException('Missing required field createdBy for ArtifactVersion.'); }
        if (!array_key_exists('createdAt', $data)) { throw new InvalidArgumentException('Missing required field createdAt for ArtifactVersion.'); }
        return new self($data);
    }

    public function schemaVersion(): string
    {
        return $this->data['schemaVersion'];
    }

    public function hasSchemaVersion(): bool
    {
        return array_key_exists('schemaVersion', $this->data);
    }

    public function id(): string
    {
        return $this->data['id'];
    }

    public function hasId(): bool
    {
        return array_key_exists('id', $this->data);
    }

    public function artifactId(): string
    {
        return $this->data['artifactId'];
    }

    public function hasArtifactId(): bool
    {
        return array_key_exists('artifactId', $this->data);
    }

    public function version(): int
    {
        return $this->data['version'];
    }

    public function hasVersion(): bool
    {
        return array_key_exists('version', $this->data);
    }

    /** @return array<mixed> */
    public function source(): array
    {
        return $this->data['source'];
    }

    public function hasSource(): bool
    {
        return array_key_exists('source', $this->data);
    }

    /** @return array<mixed>|null */
    public function integrity(): ?array
    {
        return $this->data['integrity'] ?? null;
    }

    public function hasIntegrity(): bool
    {
        return array_key_exists('integrity', $this->data);
    }

    /** @return array<mixed>|null */
    public function specification(): ?array
    {
        return $this->data['specification'] ?? null;
    }

    public function hasSpecification(): bool
    {
        return array_key_exists('specification', $this->data);
    }

    /** @return array<mixed> */
    public function createdBy(): array
    {
        return $this->data['createdBy'];
    }

    public function hasCreatedBy(): bool
    {
        return array_key_exists('createdBy', $this->data);
    }

    public function createdAt(): string
    {
        return $this->data['createdAt'];
    }

    public function hasCreatedAt(): bool
    {
        return array_key_exists('createdAt', $this->data);
    }

    public function note(): ?string
    {
        return $this->data['note'] ?? null;
    }

    public function hasNote(): bool
    {
        return array_key_exists('note', $this->data);
    }

    /** @return array<mixed>|null */
    public function metadata(): ?array
    {
        return $this->data['metadata'] ?? null;
    }

    public function hasMetadata(): bool
    {
        return array_key_exists('metadata', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['schemaVersion' => true, 'id' => true, 'artifactId' => true, 'version' => true, 'source' => true, 'integrity' => true, 'specification' => true, 'createdBy' => true, 'createdAt' => true, 'note' => true, 'metadata' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
