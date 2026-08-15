<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Enum;

enum ArtifactConditionKind: string
{
    case V_STATE = 'state';
    case V_ACTOR = 'actor';
    case V_ARTIFACT_EXISTS = 'artifact_exists';
    case V_ARTIFACT_VALUE = 'artifact_value';
    case V_AND = 'and';
    case V_OR = 'or';
    case V_NOT = 'not';
}
