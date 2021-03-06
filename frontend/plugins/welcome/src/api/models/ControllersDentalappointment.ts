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
/**
 * 
 * @export
 * @interface ControllersDentalappointment
 */
export interface ControllersDentalappointment {
    /**
     * 
     * @type {number}
     * @memberof ControllersDentalappointment
     */
    amount?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersDentalappointment
     */
    appointTime?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersDentalappointment
     */
    kindName?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersDentalappointment
     */
    note?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersDentalappointment
     */
    patientID?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersDentalappointment
     */
    personnelID?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersDentalappointment
     */
    price?: number;
}

export function ControllersDentalappointmentFromJSON(json: any): ControllersDentalappointment {
    return ControllersDentalappointmentFromJSONTyped(json, false);
}

export function ControllersDentalappointmentFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersDentalappointment {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'amount': !exists(json, 'Amount') ? undefined : json['Amount'],
        'appointTime': !exists(json, 'AppointTime') ? undefined : json['AppointTime'],
        'kindName': !exists(json, 'KindName') ? undefined : json['KindName'],
        'note': !exists(json, 'Note') ? undefined : json['Note'],
        'patientID': !exists(json, 'PatientID') ? undefined : json['PatientID'],
        'personnelID': !exists(json, 'PersonnelID') ? undefined : json['PersonnelID'],
        'price': !exists(json, 'Price') ? undefined : json['Price'],
    };
}

export function ControllersDentalappointmentToJSON(value?: ControllersDentalappointment | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'amount': value.amount,
        'appointTime': value.appointTime,
        'kindName': value.kindName,
        'note': value.note,
        'patientID': value.patientID,
        'personnelID': value.personnelID,
        'price': value.price,
    };
}


