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
    EntDentalappointment,
    EntDentalappointmentFromJSON,
    EntDentalappointmentFromJSONTyped,
    EntDentalappointmentToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDentalkindEdges
 */
export interface EntDentalkindEdges {
    /**
     * Dentalappointment holds the value of the Dentalappointment edge.
     * @type {Array<EntDentalappointment>}
     * @memberof EntDentalkindEdges
     */
    dentalappointment?: Array<EntDentalappointment>;
}

export function EntDentalkindEdgesFromJSON(json: any): EntDentalkindEdges {
    return EntDentalkindEdgesFromJSONTyped(json, false);
}

export function EntDentalkindEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDentalkindEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'dentalappointment': !exists(json, 'Dentalappointment') ? undefined : ((json['Dentalappointment'] as Array<any>).map(EntDentalappointmentFromJSON)),
    };
}

export function EntDentalkindEdgesToJSON(value?: EntDentalkindEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'dentalappointment': value.dentalappointment === undefined ? undefined : ((value.dentalappointment as Array<any>).map(EntDentalappointmentToJSON)),
    };
}


