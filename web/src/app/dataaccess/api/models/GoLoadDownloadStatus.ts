/* tslint:disable */
/* eslint-disable */
/**
 * api/go_load.proto
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: version not set
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * 
 * @export
 */
export const GoLoadDownloadStatus = {
    UndefinedStatus: 'UndefinedStatus',
    Pending: 'Pending',
    Downloading: 'Downloading',
    Failed: 'Failed',
    Success: 'Success'
} as const;
export type GoLoadDownloadStatus = typeof GoLoadDownloadStatus[keyof typeof GoLoadDownloadStatus];


export function instanceOfGoLoadDownloadStatus(value: any): boolean {
    for (const key in GoLoadDownloadStatus) {
        if (Object.prototype.hasOwnProperty.call(GoLoadDownloadStatus, key)) {
            if (GoLoadDownloadStatus[key as keyof typeof GoLoadDownloadStatus] === value) {
                return true;
            }
        }
    }
    return false;
}

export function GoLoadDownloadStatusFromJSON(json: any): GoLoadDownloadStatus {
    return GoLoadDownloadStatusFromJSONTyped(json, false);
}

export function GoLoadDownloadStatusFromJSONTyped(json: any, ignoreDiscriminator: boolean): GoLoadDownloadStatus {
    return json as GoLoadDownloadStatus;
}

export function GoLoadDownloadStatusToJSON(value?: GoLoadDownloadStatus | null): any {
    return value as any;
}

export function GoLoadDownloadStatusToJSONTyped(value: any, ignoreDiscriminator: boolean): GoLoadDownloadStatus {
    return value as GoLoadDownloadStatus;
}

