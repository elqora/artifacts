<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactRequirement
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('schemaVersion', $data)) { throw new InvalidArgumentException('Missing required field schemaVersion for ArtifactRequirement.'); }
        if (!array_key_exists('id', $data)) { throw new InvalidArgumentException('Missing required field id for ArtifactRequirement.'); }
        if (!array_key_exists('required', $data)) { throw new InvalidArgumentException('Missing required field required for ArtifactRequirement.'); }
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

    public function key(): ?string
    {
        return $this->data['key'] ?? null;
    }

    public function hasKey(): bool
    {
        return array_key_exists('key', $this->data);
    }

    /** @return array<mixed>|null */
    public function allowedKinds(): ?array
    {
        return $this->data['allowedKinds'] ?? null;
    }

    public function hasAllowedKinds(): bool
    {
        return array_key_exists('allowedKinds', $this->data);
    }

    /** @return array<mixed>|null */
    public function allowedValueTypes(): ?array
    {
        return $this->data['allowedValueTypes'] ?? null;
    }

    public function hasAllowedValueTypes(): bool
    {
        return array_key_exists('allowedValueTypes', $this->data);
    }

    public function minimumCount(): ?int
    {
        return $this->data['minimumCount'] ?? null;
    }

    public function hasMinimumCount(): bool
    {
        return array_key_exists('minimumCount', $this->data);
    }

    public function maximumCount(): ?int
    {
        return $this->data['maximumCount'] ?? null;
    }

    public function hasMaximumCount(): bool
    {
        return array_key_exists('maximumCount', $this->data);
    }

    public function required(): bool
    {
        return $this->data['required'];
    }

    public function hasRequired(): bool
    {
        return array_key_exists('required', $this->data);
    }

    public function specId(): ?string
    {
        return $this->data['specId'] ?? null;
    }

    public function hasSpecId(): bool
    {
        return array_key_exists('specId', $this->data);
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
        return array_diff_key($this->data, ['schemaVersion' => true, 'id' => true, 'key' => true, 'allowedKinds' => true, 'allowedValueTypes' => true, 'minimumCount' => true, 'maximumCount' => true, 'required' => true, 'specId' => true, 'metadata' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
