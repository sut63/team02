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
    EntChecksymptomsEdges,
    EntChecksymptomsEdgesFromJSON,
    EntChecksymptomsEdgesFromJSONTyped,
    EntChecksymptomsEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntChecksymptoms
 */
export interface EntChecksymptoms {
    /**
     * Date holds the value of the "date" field.
     * @type {string}
     * @memberof EntChecksymptoms
     */
    date?: string;
    /**
     * 
     * @type {EntChecksymptomsEdges}
     * @memberof EntChecksymptoms
     */
    edges?: EntChecksymptomsEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntChecksymptoms
     */
    id?: number;
}

export function EntChecksymptomsFromJSON(json: any): EntChecksymptoms {
    return EntChecksymptomsFromJSONTyped(json, false);
}

export function EntChecksymptomsFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntChecksymptoms {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'date': !exists(json, 'date') ? undefined : json['date'],
        'edges': !exists(json, 'edges') ? undefined : EntChecksymptomsEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntChecksymptomsToJSON(value?: EntChecksymptoms | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'date': value.date,
        'edges': EntChecksymptomsEdgesToJSON(value.edges),
        'id': value.id,
    };
}

