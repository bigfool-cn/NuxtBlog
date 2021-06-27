<template>
  <div>
    <div class="label-title">时间转时间戳:</div>
    <div class="date-time">
      <el-input v-model="datetime1" placeholder="时间" />
      <span>></span>
      <el-select v-model="unit1" placeholder="请选择">
        <el-option value="s" label="秒"/>
        <el-option value="ms" label="毫秒"/>
      </el-select>
      <el-input v-model="unix1" placeholder="时间戳" />
    </div>
    <div class="label-title">时间戳转时间:</div>
    <div class="date-time">
      <el-input v-model="unix" placeholder="时间戳" />
      <el-select v-model="unit" placeholder="请选择">
        <el-option value="s" label="秒"/>
        <el-option value="ms" label="毫秒"/>
      </el-select>
      <span>></span>
      <el-input v-model="datetime" placeholder="时间" />
    </div>
    <div class="label-title">计算时间差:</div>
    <div class="datetime-diff">
      <el-input v-model="startTime" placeholder="开始时间" />
      <span>-</span>
      <el-input v-model="endTime" placeholder="结束时间" />
    </div>
    <ul>
      <li>月：{{diffMonth}}</li>
      <li>星期：{{diffWeek}}</li>
      <li>天：{{diffDay}}</li>
      <li>小时：{{diffHour}}</li>
      <li>分钟：{{diffMintue}}</li>
      <li>秒：{{diffSecond}}</li>
      <li>毫秒：{{diffMillisecond}}</li>
    </ul>
  </div>
</template>

<script>
  import * as dayjs from 'dayjs'
  export default {
    name: "ToolDatetime",
    data() {
      return {
        unix: '',
        unit: 's',
        datetime: '',
        unix1: '',
        unit1: 's',
        datetime1: '',
        startTime: '',
        endTime: ''
      }
    },
    mounted() {
      const date = new Date()
      this.unix = this.unix1 = dayjs(date).unix()
      this.datetime = this.datetime1 = this.startTime = this.endTime = dayjs(date).format('YYYY-MM-DD HH:mm:ss')
    },
    watch: {
      unix(val) {
        switch (this.unit) {
          case 's':
            this.datetime = dayjs.unix(val).format('YYYY-MM-DD HH:mm:ss')
            break
          case 'ms':
            this.datetime = dayjs(val * 1).format('YYYY-MM-DD HH:mm:ss:SSS')
            break
        }
      },
      unit(val) {
        switch (val) {
          case 's':
            this.datetime = dayjs.unix(this.unix).format('YYYY-MM-DD HH:mm:ss')
            break
          case 'ms':
            this.datetime = dayjs(this.unix * 1).format('YYYY-MM-DD HH:mm:ss:SSS')
            break
        }
      },
      datetime1(val) {
        switch (this.unit1) {
          case 's':
            this.unix1 = dayjs(val).unix()
            break
          case 'ms':
            this.unix1 = dayjs(val).valueOf()
            break
        }
      },
      unit1(val) {
        switch (val) {
          case 's':
            this.unix1 = dayjs(this.datetime1).unix()
            break
          case 'ms':
            this.unix1 = dayjs(this.datetime1).valueOf()
            break
        }
      }
    },
    computed: {
      diffMonth() {
        const start = dayjs(this.startTime)
        const end = dayjs(this.endTime)
        return end.diff(start,'M')
      },
      diffWeek() {
        const start = dayjs(this.startTime)
        const end = dayjs(this.endTime)
        return end.diff(start,'w')
      },
      diffDay() {
        const start = dayjs(this.startTime)
        const end = dayjs(this.endTime)
        return end.diff(start,'d')
      },
      diffHour() {
        const start = dayjs(this.startTime)
        const end = dayjs(this.endTime)
        return end.diff(start,'h')
      },
      diffMintue() {
        const start = dayjs(this.startTime)
        const end = dayjs(this.endTime)
        return end.diff(start,'m')
      },
      diffSecond() {
        const start = dayjs(this.startTime)
        const end = dayjs(this.endTime)
        return end.diff(start,'s')
      },
      diffMillisecond()
      {
        const start = dayjs(this.startTime)
        const end = dayjs(this.endTime)
        return end.diff(start,'ms')
      }
    }
  }
</script>

<style lang="scss" scoped>
  .label-title {
    margin: 35px 0 3px 0;
    color: var(--color);
  }
  .label-title:nth-of-type(1) {
    margin: 10px 0 3px 0;
  }
  .date-time {
    display: flex;
    flex-direction: row;
    justify-content: space-around;
    justify-items: center;
    align-items: center;
    span {
      padding: 0 10px;
      color: var(--color);
    }
  }
  .datetime-diff {
    display: flex;
    flex-direction: row;
    justify-content: space-around;
    justify-items: center;
    align-items: center;
    span {
      padding: 0 10px;
      color: var(--color);
    }
  }

  ul {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    list-style: none;
    margin: 0;
    padding: 3px 0 20px 0;
    color: var(--color);
  }

  li {
    padding-bottom: 15px;
  }

  li > span {
    font-weight: 600;
  }
</style>
