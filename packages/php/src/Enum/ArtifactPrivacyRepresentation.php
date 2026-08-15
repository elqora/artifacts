<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Enum;

enum ArtifactPrivacyRepresentation: string
{
    case V_HIDDEN = 'hidden';
    case V_MASKED = 'masked';
    case V_APPROXIMATE = 'approximate';
    case V_FULL = 'full';
}
