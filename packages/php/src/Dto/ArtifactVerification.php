<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactVerification
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('schemaVersion', $data)) { throw new InvalidArgumentException('Missing required field schemaVersion for ArtifactVerification.'); }
        if (!array_key_exists('id', $data)) { throw new InvalidArgumentException('Missing required field id for ArtifactVerification.'); }
        if (!array_key_exists('artifactId', $data)) { throw new InvalidArgumentException('Missing required field artifactId for ArtifactVerification.'); }
        if (!array_key_exists('status', $data)) { throw new InvalidArgumentException('Missing required field status for ArtifactVerification.'); }
        if (!array_key_exists('createdAt', $data)) { throw new InvalidArgumentException('Missing required field createdAt for ArtifactVerification.'); }
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

    public function submissionId(): ?string
    {
        return $this->data['submissionId'] ?? null;
    }

    public function hasSubmissionId(): bool
    {
        return array_key_exists('submissionId', $this->data);
    }

    public function status(): string
    {
        return $this->data['status'];
    }

    public function hasStatus(): bool
    {
        return array_key_exists('status', $this->data);
    }

    public function method(): ?string
    {
        return $this->data['method'] ?? null;
    }

    public function hasMethod(): bool
    {
        return array_key_exists('method', $this->data);
    }

    /** @return array<mixed>|null */
    public function verifiedBy(): ?array
    {
        return $this->data['verifiedBy'] ?? null;
    }

    public function hasVerifiedBy(): bool
    {
        return array_key_exists('verifiedBy', $this->data);
    }

    public function reason(): ?string
    {
        return $this->data['reason'] ?? null;
    }

    public function hasReason(): bool
    {
        return array_key_exists('reason', $this->data);
    }

    public function createdAt(): string
    {
        return $this->data['createdAt'];
    }

    public function hasCreatedAt(): bool
    {
        return array_key_exists('createdAt', $this->data);
    }

    public function verifiedAt(): ?string
    {
        return $this->data['verifiedAt'] ?? null;
    }

    public function hasVerifiedAt(): bool
    {
        return array_key_exists('verifiedAt', $this->data);
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
        return array_diff_key($this->data, ['schemaVersion' => true, 'id' => true, 'artifactId' => true, 'artifactVersionId' => true, 'submissionId' => true, 'status' => true, 'method' => true, 'verifiedBy' => true, 'reason' => true, 'createdAt' => true, 'verifiedAt' => true, 'metadata' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
