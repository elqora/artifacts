<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class Artifact
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('schemaVersion', $data)) { throw new InvalidArgumentException('Missing required field schemaVersion for Artifact.'); }
        if (!array_key_exists('id', $data)) { throw new InvalidArgumentException('Missing required field id for Artifact.'); }
        if (!array_key_exists('specId', $data)) { throw new InvalidArgumentException('Missing required field specId for Artifact.'); }
        if (!array_key_exists('kind', $data)) { throw new InvalidArgumentException('Missing required field kind for Artifact.'); }
        if (!array_key_exists('valueType', $data)) { throw new InvalidArgumentException('Missing required field valueType for Artifact.'); }
        if (!array_key_exists('createdBy', $data)) { throw new InvalidArgumentException('Missing required field createdBy for Artifact.'); }
        if (!array_key_exists('createdAt', $data)) { throw new InvalidArgumentException('Missing required field createdAt for Artifact.'); }
        if (!array_key_exists('updatedAt', $data)) { throw new InvalidArgumentException('Missing required field updatedAt for Artifact.'); }
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

    public function specId(): string
    {
        return $this->data['specId'];
    }

    public function hasSpecId(): bool
    {
        return array_key_exists('specId', $this->data);
    }

    /** @return array<mixed>|null */
    public function scope(): ?array
    {
        return $this->data['scope'] ?? null;
    }

    public function hasScope(): bool
    {
        return array_key_exists('scope', $this->data);
    }

    public function kind(): string
    {
        return $this->data['kind'];
    }

    public function hasKind(): bool
    {
        return array_key_exists('kind', $this->data);
    }

    public function valueType(): string
    {
        return $this->data['valueType'];
    }

    public function hasValueType(): bool
    {
        return array_key_exists('valueType', $this->data);
    }

    public function title(): ?string
    {
        return $this->data['title'] ?? null;
    }

    public function hasTitle(): bool
    {
        return array_key_exists('title', $this->data);
    }

    public function description(): ?string
    {
        return $this->data['description'] ?? null;
    }

    public function hasDescription(): bool
    {
        return array_key_exists('description', $this->data);
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

    public function currentVersionId(): ?string
    {
        return $this->data['currentVersionId'] ?? null;
    }

    public function hasCurrentVersionId(): bool
    {
        return array_key_exists('currentVersionId', $this->data);
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

    public function updatedAt(): string
    {
        return $this->data['updatedAt'];
    }

    public function hasUpdatedAt(): bool
    {
        return array_key_exists('updatedAt', $this->data);
    }

    public function archivedAt(): ?string
    {
        return $this->data['archivedAt'] ?? null;
    }

    public function hasArchivedAt(): bool
    {
        return array_key_exists('archivedAt', $this->data);
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
        return array_diff_key($this->data, ['schemaVersion' => true, 'id' => true, 'specId' => true, 'scope' => true, 'kind' => true, 'valueType' => true, 'title' => true, 'description' => true, 'specification' => true, 'currentVersionId' => true, 'createdBy' => true, 'createdAt' => true, 'updatedAt' => true, 'archivedAt' => true, 'metadata' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
