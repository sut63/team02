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
    EntBonediseaseEdges,
    EntBonediseaseEdgesFromJSON,
    EntBonediseaseEdgesFromJSONTyped,
    EntBonediseaseEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntBonedisease
 */
export interface EntBonedisease {
    /**
     * AddedTime holds the value of the "addedTime" field.
     * @type {string}
     * @memberof EntBonedisease
     */
    addedTime?: string;
    /**
     * Advice holds the value of the "advice" field.
     * @type {string}
     * @memberof EntBonedisease
     */
    advice?: string;
    /**
     * 
     * @type {EntBonediseaseEdges}
     * @memberof EntBonedisease
     */
    edges?: EntBonediseaseEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntBonedisease
     */
    id?: number;
}

export function EntBonediseaseFromJSON(json: any): EntBonedisease {
    return EntBonediseaseFromJSONTyped(json, false);
}

export function EntBonediseaseFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntBonedisease {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'addedTime': !exists(json, 'addedTime') ? undefined : json['addedTime'],
        'advice': !exists(json, 'advice') ? undefined : json['advice'],
        'edges': !exists(json, 'edges') ? undefined : EntBonediseaseEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntBonediseaseToJSON(value?: EntBonedisease | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'addedTime': value.addedTime,
        'advice': value.advice,
        'edges': EntBonediseaseEdgesToJSON(value.edges),
        'id': value.id,
    };
}

