<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Enum;

enum ArtifactPrivacyClassification: string
{
    case V_PUBLIC = 'public';
    case V_INTERNAL = 'internal';
    case V_PRIVATE = 'private';
    case V_SENSITIVE = 'sensitive';
    case V_RESTRICTED = 'restricted';
}
