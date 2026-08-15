<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactPresentationPolicy
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        // This schema has no required fields.
        return new self($data);
    }

    public function label(): ?string
    {
        return $this->data['label'] ?? null;
    }

    public function hasLabel(): bool
    {
        return array_key_exists('label', $this->data);
    }

    public function helpText(): ?string
    {
        return $this->data['helpText'] ?? null;
    }

    public function hasHelpText(): bool
    {
        return array_key_exists('helpText', $this->data);
    }

    public function order(): float|int|null
    {
        return $this->data['order'] ?? null;
    }

    public function hasOrder(): bool
    {
        return array_key_exists('order', $this->data);
    }

    public function display(): ?string
    {
        return $this->data['display'] ?? null;
    }

    public function hasDisplay(): bool
    {
        return array_key_exists('display', $this->data);
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
        return array_diff_key($this->data, ['label' => true, 'helpText' => true, 'order' => true, 'display' => true, 'config' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
