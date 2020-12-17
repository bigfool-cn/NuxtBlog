import Vue from 'vue'
import locale from 'element-ui/lib/locale/lang/zh-CN'

import {
  Table,
  TableColumn,
  Button,
  Form,
  FormItem,
  Input,
  Select,
  Option,
  Dialog,
  Pagination,
  Message,
  DatePicker
} from 'element-ui'

const components = [
  Table,
  TableColumn,
  Button,
  Form,
  FormItem,
  Input,
  Select,
  Option,
  Dialog,
  Pagination,
  Message,
  DatePicker
];

const Element = {
  install (Vue) {
    components.forEach(component => {
      Vue.component(component.name, component)
    })
  }
}
Vue.use(Element, { locale })
