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
    EntAntenatalinformation,
    EntAntenatalinformationFromJSON,
    EntAntenatalinformationFromJSONTyped,
    EntAntenatalinformationToJSON,
    EntBonedisease,
    EntBonediseaseFromJSON,
    EntBonediseaseFromJSONTyped,
    EntBonediseaseToJSON,
    EntChecksymptom,
    EntChecksymptomFromJSON,
    EntChecksymptomFromJSONTyped,
    EntChecksymptomToJSON,
    EntDentalappointment,
    EntDentalappointmentFromJSON,
    EntDentalappointmentFromJSONTyped,
    EntDentalappointmentToJSON,
    EntGender,
    EntGenderFromJSON,
    EntGenderFromJSONTyped,
    EntGenderToJSON,
    EntPhysicaltherapyrecord,
    EntPhysicaltherapyrecordFromJSON,
    EntPhysicaltherapyrecordFromJSONTyped,
    EntPhysicaltherapyrecordToJSON,
    EntSurgeryappointment,
    EntSurgeryappointmentFromJSON,
    EntSurgeryappointmentFromJSONTyped,
    EntSurgeryappointmentToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPatientEdges
 */
export interface EntPatientEdges {
    /**
     * Antenatalinformation holds the value of the Antenatalinformation edge.
     * @type {Array<EntAntenatalinformation>}
     * @memberof EntPatientEdges
     */
    antenatalinformation?: Array<EntAntenatalinformation>;
    /**
     * Bonedisease holds the value of the Bonedisease edge.
     * @type {Array<EntBonedisease>}
     * @memberof EntPatientEdges
     */
    bonedisease?: Array<EntBonedisease>;
    /**
     * Checksymptom holds the value of the Checksymptom edge.
     * @type {Array<EntChecksymptom>}
     * @memberof EntPatientEdges
     */
    checksymptom?: Array<EntChecksymptom>;
    /**
     * Dentalappointment holds the value of the Dentalappointment edge.
     * @type {Array<EntDentalappointment>}
     * @memberof EntPatientEdges
     */
    dentalappointment?: Array<EntDentalappointment>;
    /**
     * 
     * @type {EntGender}
     * @memberof EntPatientEdges
     */
    gender?: EntGender;
    /**
     * Physicaltherapyrecord holds the value of the physicaltherapyrecord edge.
     * @type {Array<EntPhysicaltherapyrecord>}
     * @memberof EntPatientEdges
     */
    physicaltherapyrecord?: Array<EntPhysicaltherapyrecord>;
    /**
     * Surgeryappointment holds the value of the Surgeryappointment edge.
     * @type {Array<EntSurgeryappointment>}
     * @memberof EntPatientEdges
     */
    surgeryappointment?: Array<EntSurgeryappointment>;
}

export function EntPatientEdgesFromJSON(json: any): EntPatientEdges {
    return EntPatientEdgesFromJSONTyped(json, false);
}

export function EntPatientEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPatientEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'antenatalinformation': !exists(json, 'antenatalinformation') ? undefined : ((json['antenatalinformation'] as Array<any>).map(EntAntenatalinformationFromJSON)),
        'bonedisease': !exists(json, 'bonedisease') ? undefined : ((json['bonedisease'] as Array<any>).map(EntBonediseaseFromJSON)),
        'checksymptom': !exists(json, 'checksymptom') ? undefined : ((json['checksymptom'] as Array<any>).map(EntChecksymptomFromJSON)),
        'dentalappointment': !exists(json, 'dentalappointment') ? undefined : ((json['dentalappointment'] as Array<any>).map(EntDentalappointmentFromJSON)),
        'gender': !exists(json, 'gender') ? undefined : EntGenderFromJSON(json['gender']),
        'physicaltherapyrecord': !exists(json, 'physicaltherapyrecord') ? undefined : ((json['physicaltherapyrecord'] as Array<any>).map(EntPhysicaltherapyrecordFromJSON)),
        'surgeryappointment': !exists(json, 'surgeryappointment') ? undefined : ((json['surgeryappointment'] as Array<any>).map(EntSurgeryappointmentFromJSON)),
    };
}

export function EntPatientEdgesToJSON(value?: EntPatientEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'antenatalinformation': value.antenatalinformation === undefined ? undefined : ((value.antenatalinformation as Array<any>).map(EntAntenatalinformationToJSON)),
        'bonedisease': value.bonedisease === undefined ? undefined : ((value.bonedisease as Array<any>).map(EntBonediseaseToJSON)),
        'checksymptom': value.checksymptom === undefined ? undefined : ((value.checksymptom as Array<any>).map(EntChecksymptomToJSON)),
        'dentalappointment': value.dentalappointment === undefined ? undefined : ((value.dentalappointment as Array<any>).map(EntDentalappointmentToJSON)),
        'gender': EntGenderToJSON(value.gender),
        'physicaltherapyrecord': value.physicaltherapyrecord === undefined ? undefined : ((value.physicaltherapyrecord as Array<any>).map(EntPhysicaltherapyrecordToJSON)),
        'surgeryappointment': value.surgeryappointment === undefined ? undefined : ((value.surgeryappointment as Array<any>).map(EntSurgeryappointmentToJSON)),
    };
}


