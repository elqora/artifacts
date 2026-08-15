<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactMaskingPolicy
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('strategy', $data)) { throw new InvalidArgumentException('Missing required field strategy for ArtifactMaskingPolicy.'); }
        return new self($data);
    }

    public function strategy(): string
    {
        return $this->data['strategy'];
    }

    public function hasStrategy(): bool
    {
        return array_key_exists('strategy', $this->data);
    }

    /** @return array<mixed>|null */
    public function config(): ?array
    {
        return $this->data['config'] ?? null;
    }

    public function hasConfig(): bool
    {
        return array_key_exists('config', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['strategy' => true, 'config' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
