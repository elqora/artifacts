<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactSubmissionContext
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        // This schema has no required fields.
        return new self($data);
    }

    public function latitude(): float|int|null
    {
        return $this->data['latitude'] ?? null;
    }

    public function hasLatitude(): bool
    {
        return array_key_exists('latitude', $this->data);
    }

    public function longitude(): float|int|null
    {
        return $this->data['longitude'] ?? null;
    }

    public function hasLongitude(): bool
    {
        return array_key_exists('longitude', $this->data);
    }

    public function deviceId(): ?string
    {
        return $this->data['deviceId'] ?? null;
    }

    public function hasDeviceId(): bool
    {
        return array_key_exists('deviceId', $this->data);
    }

    public function ipAddress(): ?string
    {
        return $this->data['ipAddress'] ?? null;
    }

    public function hasIpAddress(): bool
    {
        return array_key_exists('ipAddress', $this->data);
    }

    public function userAgent(): ?string
    {
        return $this->data['userAgent'] ?? null;
    }

    public function hasUserAgent(): bool
    {
        return array_key_exists('userAgent', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['latitude' => true, 'longitude' => true, 'deviceId' => true, 'ipAddress' => true, 'userAgent' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
