<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactVerificationPolicy
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('required', $data)) { throw new InvalidArgumentException('Missing required field required for ArtifactVerificationPolicy.'); }
        return new self($data);
    }

    public function required(): bool
    {
        return $this->data['required'];
    }

    public function hasRequired(): bool
    {
        return array_key_exists('required', $this->data);
    }

    /** @return array<mixed>|null */
    public function methods(): ?array
    {
        return $this->data['methods'] ?? null;
    }

    public function hasMethods(): bool
    {
        return array_key_exists('methods', $this->data);
    }

    /** @return array<mixed>|null */
    public function actors(): ?array
    {
        return $this->data['actors'] ?? null;
    }

    public function hasActors(): bool
    {
        return array_key_exists('actors', $this->data);
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
        return array_diff_key($this->data, ['required' => true, 'methods' => true, 'actors' => true, 'condition' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
