<template>
  <v-container>
    <v-row class="mb-6">
      <v-col>
        <v-btn color="primary" @click="generateYesterdaysReadings">Generate readings</v-btn>
        <v-btn color="secondary" class="ml-4" @click="calculateCost">Calculate cost</v-btn>
      </v-col>
    </v-row>

    <!-- Display the calculated cost -->
    <v-row>
      <v-col>
        <v-card outlined>
          <v-card-title>Calculated Cost</v-card-title>
          <v-card-subtitle>
            Cost: €{{ cost.toLocaleString('en', { maximumFractionDigits: 2, minimumFractionDigits: 2 }) }}
          </v-card-subtitle>
        </v-card>
      </v-col>
    </v-row>

    <!-- Table to display meter readings -->
    <v-row>
      <v-col>
        <v-data-table :headers="tableHeaders" :items="meterReadings">
          <!-- Custom formatting for timestamp -->
          <template v-slot:item.timestamp="{ item }">
            {{ new Date(item.timestamp).toLocaleString() }}
          </template>
          <!-- Custom formatting for kWh -->
          <template v-slot:item.kwh="{ item }">
            {{ item.kwh.toFixed(2) }}
          </template>
        </v-data-table>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
import { ref, type Ref } from 'vue'
import axios from 'axios'

interface MeterReading {
  timestamp: Date
  kwh: number
}

const meterReadings: Ref<MeterReading[]> = ref([])
const cost: Ref<number> = ref(0)

// Table headers for Vuetify data table
const tableHeaders = [
  { title: 'Time', key: 'timestamp' },
  { title: 'Meter Reading (kWh)', key: 'kwh' }
]

// Generates random meter readings spanning the 24 hours of yesterday's date.
function generateYesterdaysReadings() {
  var kwh = Math.random() * 1000000
  const readings = []

  for (var hour = 0; hour <= 24; hour++) {
    const date = new Date()
    date.setDate(date.getDate() - 1)
    date.setHours(hour, 0, 0, 0)

    kwh += Math.random() * 100

    readings.push({
      timestamp: date.getTime(),
      kwh: kwh
    })
  }

  meterReadings.value = readings
}

// Function to calculate energy cost by sending meter readings to backend
async function calculateCost() {
  try {
    const response = await axios.post('http://localhost:8080/energy_cost', meterReadings.value)
    cost.value = response.data.total_cost
  } catch (error) {
    console.error('Failed to calculate energy cost:', error)
  }
}
</script>

<!--<style scoped>-->
<!--.v-container {-->
<!--  padding: 30px;-->
<!--}-->
<!--</style>-->



<!--<template>-->
<!--  <div id="app">-->
<!--    <div class="btns">-->
<!--      <button @click="generateYesterdaysReadings()">Generate readings</button>-->
<!--      <button @click="calculateCost()">Calculate cost</button>-->
<!--    </div>-->

<!--    &lt;!&ndash; Display the calculated cost &ndash;&gt;-->
<!--    <div>-->
<!--      Cost: €{{ cost.toLocaleString('en', { maximumFractionDigits: 2, minimumFractionDigits: 2 }) }}-->
<!--    </div>-->

<!--    &lt;!&ndash; Table to display meter readings &ndash;&gt;-->
<!--    <table>-->
<!--      <thead>-->
<!--        <tr>-->
<!--          <th>Time</th>-->
<!--          <th>Meter Reading (kWh)</th>-->
<!--        </tr>-->
<!--      </thead>-->
<!--      <tbody>-->
<!--        <tr v-for="(reading, index) in meterReadings" :key="index">-->
<!--          <td>{{ new Date(reading.timestamp).toLocaleString() }}</td>-->
<!--          <td>{{ reading.kwh.toFixed(2) }}</td>-->
<!--        </tr>-->
<!--      </tbody>-->
<!--    </table>-->

<!--  </div>-->
<!--</template>-->

<!--<script setup lang="ts">-->
<!--import { ref, type Ref } from 'vue'-->
<!--import axios from 'axios'-->

<!--interface MeterReading {-->
<!--  timestamp: Date-->
<!--  kwh: number-->
<!--}-->

<!--const meterReadings: Ref<MeterReading[]> = ref([])-->
<!--const cost: Ref<number> = ref(0)-->

<!--// Generates random meter readings spanning the 24 hours of yesterday's date.-->
<!--function generateYesterdaysReadings() {-->
<!--  var kwh = Math.random() * 1000000-->
<!--  const readings = []-->

<!--  for (var hour = 0; hour <= 24; hour++) {-->
<!--    const date = new Date()-->
<!--    date.setDate(date.getDate() - 1)-->
<!--    date.setHours(hour, 0, 0, 0)-->

<!--    kwh += Math.random() * 100-->

<!--    readings.push({-->
<!--      timestamp: date.getTime(),-->
<!--      kwh: kwh-->
<!--    })-->
<!--  }-->

<!--  meterReadings.value = readings-->
<!--}-->

<!--// Function to calculate energy cost by sending meter readings to backend-->
<!--async function calculateCost() {-->
<!--  try {-->
<!--    const response = await axios.post('http://localhost:8080/energy_cost', meterReadings.value)-->
<!--    cost.value = response.data.total_cost-->
<!--  } catch (error) {-->
<!--    console.error('Failed to calculate energy cost:', error)-->
<!--  }-->
<!--}-->
<!--</script>-->

<!--<style scoped>-->
<!--#app {-->
<!--  padding: 30px;-->
<!--  display: flex;-->
<!--  flex-direction: column;-->
<!--  gap: 30px;-->
<!--}-->

<!--table {-->
<!--  width: 100%;-->
<!--  border-collapse: collapse;-->
<!--}-->

<!--th, td {-->
<!--  padding: 8px;-->
<!--  text-align: left;-->
<!--  border-bottom: 1px solid #ddd;-->
<!--}-->
<!--</style>-->
