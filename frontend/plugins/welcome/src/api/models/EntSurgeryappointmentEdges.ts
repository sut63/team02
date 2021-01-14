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
    EntPatient,
    EntPatientFromJSON,
    EntPatientFromJSONTyped,
    EntPatientToJSON,
    EntPersonnel,
    EntPersonnelFromJSON,
    EntPersonnelFromJSONTyped,
    EntPersonnelToJSON,
    EntSurgerytype,
    EntSurgerytypeFromJSON,
    EntSurgerytypeFromJSONTyped,
    EntSurgerytypeToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntSurgeryappointmentEdges
 */
export interface EntSurgeryappointmentEdges {
    /**
     * 
     * @type {EntPatient}
     * @memberof EntSurgeryappointmentEdges
     */
    patient?: EntPatient;
    /**
     * 
     * @type {EntPersonnel}
     * @memberof EntSurgeryappointmentEdges
     */
    personnel?: EntPersonnel;
    /**
     * 
     * @type {EntSurgerytype}
     * @memberof EntSurgeryappointmentEdges
     */
    surgerytype?: EntSurgerytype;
}

export function EntSurgeryappointmentEdgesFromJSON(json: any): EntSurgeryappointmentEdges {
    return EntSurgeryappointmentEdgesFromJSONTyped(json, false);
}

export function EntSurgeryappointmentEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntSurgeryappointmentEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'patient': !exists(json, 'patient') ? undefined : EntPatientFromJSON(json['patient']),
        'personnel': !exists(json, 'personnel') ? undefined : EntPersonnelFromJSON(json['personnel']),
        'surgerytype': !exists(json, 'surgerytype') ? undefined : EntSurgerytypeFromJSON(json['surgerytype']),
    };
}

export function EntSurgeryappointmentEdgesToJSON(value?: EntSurgeryappointmentEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'patient': EntPatientToJSON(value.patient),
        'personnel': EntPersonnelToJSON(value.personnel),
        'surgerytype': EntSurgerytypeToJSON(value.surgerytype),
    };
}


