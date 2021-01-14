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
    EntRemedyEdges,
    EntRemedyEdgesFromJSON,
    EntRemedyEdgesFromJSONTyped,
    EntRemedyEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntRemedy
 */
export interface EntRemedy {
    /**
     * 
     * @type {EntRemedyEdges}
     * @memberof EntRemedy
     */
    edges?: EntRemedyEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntRemedy
     */
    id?: number;
    /**
     * Remedy holds the value of the "remedy" field.
     * @type {string}
     * @memberof EntRemedy
     */
    remedy?: string;
}

export function EntRemedyFromJSON(json: any): EntRemedy {
    return EntRemedyFromJSONTyped(json, false);
}

export function EntRemedyFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntRemedy {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntRemedyEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'remedy': !exists(json, 'remedy') ? undefined : json['remedy'],
    };
}

export function EntRemedyToJSON(value?: EntRemedy | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntRemedyEdgesToJSON(value.edges),
        'id': value.id,
        'remedy': value.remedy,
    };
}


