<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Enum;

enum ArtifactSourceType: string
{
    case V_INLINE = 'inline';
    case V_LOCAL = 'local';
    case V_OBJECT = 'object';
    case V_URL = 'url';
    case V_HOSTED = 'hosted';
    case V_PROVIDER = 'provider';
}
