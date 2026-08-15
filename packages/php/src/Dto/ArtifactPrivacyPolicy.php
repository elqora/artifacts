<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactPrivacyPolicy
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('classification', $data)) { throw new InvalidArgumentException('Missing required field classification for ArtifactPrivacyPolicy.'); }
        return new self($data);
    }

    public function classification(): string
    {
        return $this->data['classification'];
    }

    public function hasClassification(): bool
    {
        return array_key_exists('classification', $this->data);
    }

    /** @return array<mixed>|null */
    public function reveal(): ?array
    {
        return $this->data['reveal'] ?? null;
    }

    public function hasReveal(): bool
    {
        return array_key_exists('reveal', $this->data);
    }

    /** @return array<mixed>|null */
    public function masking(): ?array
    {
        return $this->data['masking'] ?? null;
    }

    public function hasMasking(): bool
    {
        return array_key_exists('masking', $this->data);
    }

    /** @return array<mixed>|null */
    public function encryption(): ?array
    {
        return $this->data['encryption'] ?? null;
    }

    public function hasEncryption(): bool
    {
        return array_key_exists('encryption', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['classification' => true, 'reveal' => true, 'masking' => true, 'encryption' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
