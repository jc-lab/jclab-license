<template>
    <div>
        <h1>License Verification</h1>
        <table>
            <tr>
                <td>LICENSE CODE</td>
                <td><textarea v-model="licenseCode" placeholder="ey...eJ....aaa" cols="100" rows="3"></textarea></td>
            </tr>
        </table>
        <div :class="{'verified': licenseVerified, 'invalid': !licenseVerified}">{{ licenseStatus }}</div>
        <div v-if="licenseClaims">
            <table>
                <tr>
                    <td>License Issuer</td>
                    <td>{{licenseClaims.iss}}</td>
                </tr>
                <tr>
                    <td>License Product</td>
                    <td>{{licenseClaims.product}}</td>
                </tr>
                <tr>
                    <td>License Max Version</td>
                    <td>{{licenseClaims.licmaxv}} ({{ parseMaxVer(licenseClaims.licmaxv) }})</td>
                </tr>
                <tr>
                    <td>License Type</td>
                    <td>{{licenseClaims.lictype}}</td>
                </tr>
                <tr>
                    <td>Licensee Name</td>
                    <td>{{licenseClaims.lisname}}</td>
                </tr>
                <tr>
                    <td>Licensee Email</td>
                    <td>{{licenseClaims.lisemail}}</td>
                </tr>
                <tr>
                    <td>RAW</td>
                    <td><pre>{{ JSON.stringify(licenseClaims, null, 2) }}</pre></td>
                </tr>
            </table>
        </div>
    </div>
</template>
<style scoped>
.verified {
    color: #278627;
}
.invalid {
    color: #c61515;
}
</style>
<script lang="ts" setup>
import {ref, watch} from 'vue';
import {verify, Claims} from './jclab-license';

const licenseCode = ref('eyJhbGciOiJFZERTQSIsImtpZCI6Ilo5LVJ1d0RaVVBvIn0.eJwszE9LhEAcxvG3Es9ZD0YONrcQKUq6GBFe4uc04dT8w9-4q7vse1_EPT4PfD9nGEqQhXisKvFwX1YZDDMkXuu8pQEZ4hR-ZpUgweSi1chgjUpr1JAIUXsO86Ru979eR0jESH-iH5fT6GwuFm-7r2J9775rf3x-4UBNP_w25dtnfOK9c7QcIPNiG-zJbfaH5nTXGqU97zprR8ZCApdrAAAA___ZJjtb.h7wAValQU7rvdgCPaVWysgjx18g5LY2A1hnKQCAIYPLDQrmzDVlsZdLqWjOGIRhl3RsPdEGRcKNvV8_6lK9nDw');
const licenseStatus = ref('');
const licenseVerified = ref(false);
const licenseClaims = ref<Claims>();

function parseMaxVer(maxv: number): string {
  if (maxv < 0) {
    return '영구';
  }
  const date = new Date(maxv * 86400000);
  return date.toLocaleString();
}

watch(licenseCode, (licenseCode) => {
  if (licenseCode.split('.').length !== 3) {
    licenseStatus.value = 'Please enter license code';
    licenseVerified.value = false;
    licenseClaims.value = undefined;
  } else {
    verify(licenseCode)
        .then((claims) => {
          licenseStatus.value = 'License Verified';
          licenseVerified.value = true;
          licenseClaims.value = claims;
        })
        .catch((err) => {
          licenseStatus.value = `License Verification Failed: ${err}`;
          licenseVerified.value = false;
          licenseClaims.value = undefined;
        });
  }
}, { immediate: true });
</script>