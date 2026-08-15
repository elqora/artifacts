<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Enum;

enum ArtifactValueType: string
{
    case V_TEXT = 'text';
    case V_NUMBER = 'number';
    case V_BOOLEAN = 'boolean';
    case V_CURRENCY = 'currency';
    case V_DATE = 'date';
    case V_DATETIME = 'datetime';
    case V_TIME = 'time';
    case V_LOCATION = 'location';
    case V_FILE = 'file';
    case V_IMAGE = 'image';
    case V_VIDEO = 'video';
    case V_AUDIO = 'audio';
    case V_LINK = 'link';
    case V_STRUCTURED = 'structured';
    case V_REFERENCE = 'reference';
    case V_SIGNATURE = 'signature';
    case V_COLLECTION = 'collection';
}
