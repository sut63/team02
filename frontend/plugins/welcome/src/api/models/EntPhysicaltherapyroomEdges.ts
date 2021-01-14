/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntPhysicaltherapyrecord,
    EntPhysicaltherapyrecordFromJSON,
    EntPhysicaltherapyrecordFromJSONTyped,
    EntPhysicaltherapyrecordToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPhysicaltherapyroomEdges
 */
export interface EntPhysicaltherapyroomEdges {
    /**
     * 
     * @type {EntPhysicaltherapyrecord}
     * @memberof EntPhysicaltherapyroomEdges
     */
    physicaltherapyrecord?: EntPhysicaltherapyrecord;
}

export function EntPhysicaltherapyroomEdgesFromJSON(json: any): EntPhysicaltherapyroomEdges {
    return EntPhysicaltherapyroomEdgesFromJSONTyped(json, false);
}

export function EntPhysicaltherapyroomEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPhysicaltherapyroomEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'physicaltherapyrecord': !exists(json, 'physicaltherapyrecord') ? undefined : EntPhysicaltherapyrecordFromJSON(json['physicaltherapyrecord']),
    };
}

export function EntPhysicaltherapyroomEdgesToJSON(value?: EntPhysicaltherapyroomEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'physicaltherapyrecord': EntPhysicaltherapyrecordToJSON(value.physicaltherapyrecord),
    };
}


