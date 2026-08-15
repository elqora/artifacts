<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactSubmission
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('schemaVersion', $data)) { throw new InvalidArgumentException('Missing required field schemaVersion for ArtifactSubmission.'); }
        if (!array_key_exists('id', $data)) { throw new InvalidArgumentException('Missing required field id for ArtifactSubmission.'); }
        if (!array_key_exists('artifactId', $data)) { throw new InvalidArgumentException('Missing required field artifactId for ArtifactSubmission.'); }
        if (!array_key_exists('submittedBy', $data)) { throw new InvalidArgumentException('Missing required field submittedBy for ArtifactSubmission.'); }
        if (!array_key_exists('submittedAt', $data)) { throw new InvalidArgumentException('Missing required field submittedAt for ArtifactSubmission.'); }
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
    public function submittedBy(): array
    {
        return $this->data['submittedBy'];
    }

    public function hasSubmittedBy(): bool
    {
        return array_key_exists('submittedBy', $this->data);
    }

    public function value(): mixed
    {
        return $this->data['value'] ?? null;
    }

    public function hasValue(): bool
    {
        return array_key_exists('value', $this->data);
    }

    public function submittedAt(): string
    {
        return $this->data['submittedAt'];
    }

    public function hasSubmittedAt(): bool
    {
        return array_key_exists('submittedAt', $this->data);
    }

    /** @return array<mixed>|null */
    public function context(): ?array
    {
        return $this->data['context'] ?? null;
    }

    public function hasContext(): bool
    {
        return array_key_exists('context', $this->data);
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
        return array_diff_key($this->data, ['schemaVersion' => true, 'id' => true, 'artifactId' => true, 'artifactVersionId' => true, 'submittedBy' => true, 'value' => true, 'submittedAt' => true, 'context' => true, 'metadata' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
