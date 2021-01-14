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
    EntPregnancystatus,
    EntPregnancystatusFromJSON,
    EntPregnancystatusFromJSONTyped,
    EntPregnancystatusToJSON,
    EntRisks,
    EntRisksFromJSON,
    EntRisksFromJSONTyped,
    EntRisksToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntAntenatalinformationEdges
 */
export interface EntAntenatalinformationEdges {
    /**
     * 
     * @type {EntPatient}
     * @memberof EntAntenatalinformationEdges
     */
    patient?: EntPatient;
    /**
     * 
     * @type {EntPersonnel}
     * @memberof EntAntenatalinformationEdges
     */
    personnel?: EntPersonnel;
    /**
     * 
     * @type {EntPregnancystatus}
     * @memberof EntAntenatalinformationEdges
     */
    pregnancystatus?: EntPregnancystatus;
    /**
     * 
     * @type {EntRisks}
     * @memberof EntAntenatalinformationEdges
     */
    risks?: EntRisks;
}

export function EntAntenatalinformationEdgesFromJSON(json: any): EntAntenatalinformationEdges {
    return EntAntenatalinformationEdgesFromJSONTyped(json, false);
}

export function EntAntenatalinformationEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntAntenatalinformationEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'patient': !exists(json, 'patient') ? undefined : EntPatientFromJSON(json['patient']),
        'personnel': !exists(json, 'personnel') ? undefined : EntPersonnelFromJSON(json['personnel']),
        'pregnancystatus': !exists(json, 'pregnancystatus') ? undefined : EntPregnancystatusFromJSON(json['pregnancystatus']),
        'risks': !exists(json, 'risks') ? undefined : EntRisksFromJSON(json['risks']),
    };
}

export function EntAntenatalinformationEdgesToJSON(value?: EntAntenatalinformationEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'patient': EntPatientToJSON(value.patient),
        'personnel': EntPersonnelToJSON(value.personnel),
        'pregnancystatus': EntPregnancystatusToJSON(value.pregnancystatus),
        'risks': EntRisksToJSON(value.risks),
    };
}


