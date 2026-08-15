<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactLink
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('schemaVersion', $data)) { throw new InvalidArgumentException('Missing required field schemaVersion for ArtifactLink.'); }
        if (!array_key_exists('id', $data)) { throw new InvalidArgumentException('Missing required field id for ArtifactLink.'); }
        if (!array_key_exists('artifactId', $data)) { throw new InvalidArgumentException('Missing required field artifactId for ArtifactLink.'); }
        if (!array_key_exists('subject', $data)) { throw new InvalidArgumentException('Missing required field subject for ArtifactLink.'); }
        if (!array_key_exists('role', $data)) { throw new InvalidArgumentException('Missing required field role for ArtifactLink.'); }
        if (!array_key_exists('createdBy', $data)) { throw new InvalidArgumentException('Missing required field createdBy for ArtifactLink.'); }
        if (!array_key_exists('createdAt', $data)) { throw new InvalidArgumentException('Missing required field createdAt for ArtifactLink.'); }
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

    public function artifactVersionId(): ?string
    {
        return $this->data['artifactVersionId'] ?? null;
    }

    public function hasArtifactVersionId(): bool
    {
        return array_key_exists('artifactVersionId', $this->data);
    }

    /** @return array<mixed> */
    public function subject(): array
    {
        return $this->data['subject'];
    }

    public function hasSubject(): bool
    {
        return array_key_exists('subject', $this->data);
    }

    public function role(): string
    {
        return $this->data['role'];
    }

    public function hasRole(): bool
    {
        return array_key_exists('role', $this->data);
    }

    public function note(): ?string
    {
        return $this->data['note'] ?? null;
    }

    public function hasNote(): bool
    {
        return array_key_exists('note', $this->data);
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
        return array_diff_key($this->data, ['schemaVersion' => true, 'id' => true, 'artifactId' => true, 'artifactVersionId' => true, 'subject' => true, 'role' => true, 'note' => true, 'createdBy' => true, 'createdAt' => true, 'metadata' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
