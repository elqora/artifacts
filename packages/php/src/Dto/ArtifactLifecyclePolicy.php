<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactLifecyclePolicy
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        // This schema has no required fields.
        return new self($data);
    }

    public function createAt(): ?string
    {
        return $this->data['createAt'] ?? null;
    }

    public function hasCreateAt(): bool
    {
        return array_key_exists('createAt', $this->data);
    }

    /** @return array<mixed>|null */
    public function editableDuring(): ?array
    {
        return $this->data['editableDuring'] ?? null;
    }

    public function hasEditableDuring(): bool
    {
        return array_key_exists('editableDuring', $this->data);
    }

    /** @return array<mixed>|null */
    public function submitDuring(): ?array
    {
        return $this->data['submitDuring'] ?? null;
    }

    public function hasSubmitDuring(): bool
    {
        return array_key_exists('submitDuring', $this->data);
    }

    public function lockAt(): ?string
    {
        return $this->data['lockAt'] ?? null;
    }

    public function hasLockAt(): bool
    {
        return array_key_exists('lockAt', $this->data);
    }

    /** @return array<mixed>|null */
    public function invalidateOn(): ?array
    {
        return $this->data['invalidateOn'] ?? null;
    }

    public function hasInvalidateOn(): bool
    {
        return array_key_exists('invalidateOn', $this->data);
    }

    /** @return array<mixed>|null */
    public function condition(): ?array
    {
        return $this->data['condition'] ?? null;
    }

    public function hasCondition(): bool
    {
        return array_key_exists('condition', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['createAt' => true, 'editableDuring' => true, 'submitDuring' => true, 'lockAt' => true, 'invalidateOn' => true, 'condition' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
