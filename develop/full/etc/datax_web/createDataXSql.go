package main

import (
	"fmt"
	"os"
	"time"
)

var tableList = []string{
	"acc_pingzheng",
	"acc_pingzheng_item",
	"acc_pingzheng_zi",
	"acc_summary",
	"acc_summary_lab",
	"acc_voucher",
	"acc_voucher_lab",
	"act_nav_menu",
	"act_sensitive",
	"action_list",
	"appupdate",
	"baojia",
	"baojia_guding_fee",
	"baojia_items",
	"baoxiao",
	"baoxiao_items",
	"baoxiao_te",
	"bom_temp",
	"bom_ver",
	"bom_ver_items",
	"cailiao_lab",
	"cangku",
	"cangku_record",
	"chu_chai_shen_qing",
	"ck_baofei_items",
	"ck_dai_chu_ru_dan",
	"ck_dan_category",
	"ck_diao_bo_items",
	"ck_faliao_items",
	"ck_jieyong_items",
	"ck_jieyong_items_return",
	"ck_op_chuhuo_items",
	"ck_op_chuhuo_items_modify",
	"ck_orders_chutui_dan",
	"ck_purchase_items_ru_tui",
	"ck_purchase_items_ru_tui_modify",
	"ck_purchase_ruku_dan",
	"ck_qita_churuku_dan",
	"ck_sc_bu_liang_items",
	"ck_sc_cheng_pin_items",
	"ck_sc_ll_items",
	"ck_sc_yuan_cai_dan",
	"ck_sc_yuan_cai_ll_items",
	"ck_zhan_yong",
	"ck_zhuan_huan_items",
	"common_shenpi",
	"company_account",
	"content_used_save",
	"customer_and_gys",
	"customer_and_gys_pingshen",
	"customer_category_list",
	"customer_category_tree",
	"customer_price",
	"daily_report",
	"daily_report_item",
	"daily_report_project",
	"department_job",
	"department_list",
	"department_tree",
	"docs",
	"docs_category_list",
	"email_record",
	"fa_piao",
	"fa_piao_items",
	"fi_bank_account",
	"fi_caigou_duizhang",
	"fi_caigou_duizhang_items",
	"fi_fukuan_dan",
	"fi_fukuan_dan_items",
	"fi_jiekuan",
	"fi_order_duizhang",
	"fi_order_duizhang_items",
	"fi_order_duizhang_items_copy1",
	"fi_qita_shou_fu",
	"fi_shiji_bank_record",
	"fi_shiji_kaipiao",
	"fi_shou_zhi_category",
	"fi_shou_zhi_record",
	"fi_shoukuan_dan",
	"fi_shoukuan_dan_items",
	"file_category_list",
	"file_category_tree",
	"file_cover",
	"file_ping_shen",
	"gantt_link",
	"gongdan_report",
	"gongduan",
	"gongxu",
	"gongyi",
	"gp_danwei",
	"gp_danwei_items",
	"guding_zichan",
	"gx_bad_items",
	"hege_peng_qi_guige",
	"hegezheng",
	"hr_jiaban_itmes",
	"hr_jiabandan",
	"id_record",
	"id_record_prefix",
	"jiao_jie_dan",
	"kaoqin_device",
	"kaoqin_recordlog",
	"kaoqin_sync_time",
	"mrp_dan",
	"mrp_dan_items",
	"mulu_category_list",
	"mulu_jihe",
	"op_kegongliao",
	"operation_doc_settings",
	"operation_log",
	"order_catalogue",
	"order_category",
	"order_file_category",
	"order_pingshen",
	"order_pingshen_items",
	"order_sn_prefix",
	"order_yan_shou_dan",
	"order_yewu_category",
	"orders",
	"orders_products",
	"orders_products_jihua",
	"orders_products_sc_bom",
	"orders_products_sc_bom_list",
	"page_content_set",
	"page_folder",
	"page_head_set",
	"plc",
	"printer",
	"procurement_req",
	"procurement_req_items",
	"product",
	"product_",
	"product_attribute",
	"product_attribute_options",
	"product_attribute_value",
	"product_attribute_verrules",
	"product_canshu",
	"product_category_list",
	"product_category_tree",
	"product_for_sell",
	"product_sku",
	"product_zxb_fee",
	"project",
	"project_catalogue",
	"project_category_list",
	"project_task",
	"psn_prefix",
	"purchase_category",
	"qc_reports",
	"qing_gou",
	"qing_gou_items",
	"sc_baogong",
	"sc_baogong_bl_item",
	"sc_cailiao",
	"sc_cailiao_ling_yong",
	"sc_chan_cheng_pin",
	"sc_features_template",
	"sc_gongxu",
	"sc_gongxu_pai_gong",
	"sc_guan_hao",
	"sc_gx_change",
	"sc_jihua_items",
	"sc_num_list",
	"sc_position",
	"sc_products",
	"sc_products_pai_gong",
	"sc_products_visual",
	"sc_purchase",
	"sc_purchase_items",
	"sc_qianshou",
	"sc_zaizhi_gx_filed",
	"sc_zhan_jie_cheng_bao",
	"sc_zhuan_chan_dan",
	"sc_zhuanchu",
	"scgx_1_xia_liao",
	"scgx_1_xia_liao_canshu",
	"scgx_2_jiao_xing",
	"scgx_3_ji_jia",
	"scgx_4_ji_guang",
	"scgx_5_ban_cheng_pin",
	"scgx_6_tui_huo",
	"scgx_6_tui_huo_canshu",
	"scgx_7_qing_xi",
	"scgx_8_hong_gan",
	"scgx_8_hong_gan_canshu",
	"scgx_9_han_jie",
	"scgx_10_cheng_zhong_1",
	"scgx_11_zi_dong",
	"scgx_11_zi_dong_canshu",
	"scgx_12_cheng_zhong_2",
	"scgx_13_feng_kou",
	"scgx_14_cheng_zhong_3",
	"scgx_15_zhe_wan",
	"scgx_16_nai_wen",
	"scgx_16_nai_wen_canshu",
	"scgx_17_cheng_zhong_4",
	"scgx_18_bao_shi_zhi",
	"scgx_19_deng_wen",
	"scgx_19_deng_wen_canshu",
	"scgx_20_dian_jiao",
	"scgx_21_zhong_jian",
	"scgx_22_zhuang_xiang",
	"scgx_label",
	"set_cangku",
	"set_cangwei",
	"shebei",
	"shebei_category_list",
	"shebei_fix",
	"shenpi_amount",
	"shenpi_dan",
	"shenpi_ren",
	"shenpi_ren_set",
	"shenpi_rst",
	"shenpi_type",
	"shenpi_type_count",
	"shenpi_type_yushe_num",
	"staff_leave",
	"staff_vacation",
	"supplier_category_list",
	"supplier_category_tree",
	"table_column",
	"temp",
	"tongzhi_setting",
	"upload_file",
	"upload_file_used",
	"upload_img",
	"user_home_kuai_jie",
	"user_role_list",
	"users",
	"users_job_level",
	"wx_temp_msg",
	"xiang_bian_cai_liao",
	"yi_chang_category",
	"yi_chang_dan",
	"yi_chang_items",
	"yi_chang_report",
	"yi_chang_report_resolve",
	"yt_content",
	"yuan_cai_options",
	"zdy_beizhu",
	"zdy_pro_field",
	"zdy_title",
	"zhuan_an",
	"zhuan_an_items",
	"zhuan_an_xuqiu",
	"zhuan_an_xuqiu_items",
}

var defaultReaderColumn = `*`
var readerTableColumn = map[string]string{
	"product_category_list": "id,label,company_id,add_user,add_time,num,sort,level,beizhu,add_uid,ck_num,sn_prefix,sys_cate,gg_prefix,pid,updated_at",
}

var defaultWriteColumn = `[\"*\"]`

// 表字段包含mysql关键词，需要转移一下，否则datax不支持
var writeTableColumn = map[string]string{
	"email_record":          `[\"id\",\"company_id\",\"subject\",\"body\",\"to_adds\",\"` + "`from`" + `\",\"cc_adds\",\"result\",\"add_time\",\"sType\",\"updated_at\"]`,
	"company_account":       `[\"id\",\"name\",\"full_name\",\"en_name\",\"account_num\",\"account_limit\",\"status\",\"logo_img\",\"word_logo\",\"file_size\",\"file_count\",\"action_list\",\"` + "`sensitive`" + `\",\"menu\",\"beizhu\",\"is_admin\",\"file_size_limit\",\"login_name\",\"up_img_size_sum\",\"up_img_count\",\"mail_pwd\",\"mail_port\",\"mail_host\",\"mail_user\",\"kq_account\",\"kq_secretkey\",\"phone\",\"adds\",\"bank_name\",\"bank_account\",\"tax_num\",\"AppID\",\"AppSecret\",\"xcx_qr_img\",\"shen_pi_set\",\"gzh_appid\",\"gzh_secret\",\"corpid\",\"agentid\",\"corpsecret\",\"gzh_qr_img\",\"bg_zhuan_qian\",\"bg_wan_zhuan\",\"bg_bad_hand\",\"first_auto_qian\",\"tmp_msg1\",\"tmp_msg2\",\"tmp_msg3\",\"tmp_msg_type\",\"book_url\",\"tmp_msg4\",\"tmp_msg5\",\"updated_at\"]`,
	"shenpi_ren_set":        `[\"id\",\"company_id\",\"copy\",\"shenpi\",\"sp_type\",\"title\",\"editable\",\"sort_index\",\"up_time\",\"beizhu\",\"` + "`default`" + `\",\"up_user\",\"updated_at\"]`,
	"upload_file":           `[\"id\",\"file_name\",\"path\",\"create_time\",\"uid\",\"company_id\",\"size\",\"ext\",\"file_sn\",\"add_user\",\"beizhu\",\"cate_id\",\"used_count\",\"` + "`where`" + `\",\"updated_at\"]`,
	"bom_ver":               `[\"id\",\"add_user\",\"add_uid\",\"add_time\",\"company_id\",\"ver\",\"p_id\",\"p_sn\",\"name\",\"guige\",\"sp_dan_sn\",\"sp_dan_id\",\"tuhao\",\"` + "`default`" + `\",\"gongxu\",\"gy_id\",\"gy_title\",\"shenhe\",\"sh_user\",\"sh_time\",\"up_time\",\"up_user\",\"updated_at\"]`,
	"ck_dai_chu_ru_dan":     `[\"id\",\"dan_sn\",\"add_user\",\"add_uid\",\"add_time\",\"company_id\",\"ru_type\",\"beizhu\",\"title_id\",\"list\",\"do_ruku\",\"ru_time\",\"ru_user\",\"ru_uid\",\"io_type\",\"extra_form_data\",\"customer_id\",\"wl_fee\",\"wl_type\",\"wl_dan_sn\",\"shou_ren\",\"shou_uid\",\"shou_phone\",\"shou_adds\",\"from_dan_sn\",\"dp_id\",\"dp_name\",\"yueding\",\"sp_dan_sn\",\"sp_dan_id\",\"kehu_ruku_status\",\"file_num\",\"yw_type\",\"shenhe\",\"sh_user\",\"sh_time\",\"c_name\",\"sc_buling\",\"closed\",\"chuType\",\"items_name\",\"items_num\",\"totalNum\",\"llDanId\",\"fenBaoid\",\"zjbid\",\"chuHuoDanId\",\"ban_cheng_bao_id\",\"mingxiList\",\"up_time\",\"up_user\",\"guan\",\"` + "`describe`" + `\",\"jf_files\",\"company\",\"updated_at\"]`,
	"operation_log":         `[\"id\",\"add_user\",\"add_uid\",\"add_time\",\"company_id\",\"method\",\"obj_id\",\"obj_sn\",\"ip_adds\",\"` + "`table`" + `\",\"beizhu\",\"updated_at\"]`,
	"file_cover":            `[\"id\",\"add_user\",\"add_uid\",\"add_time\",\"company_id\",\"title\",\"beizhu\",\"fid\",\"` + "`replace`" + `\",\"updated_at\"]`,
	"user_role_list":        `[\"id\",\"role_name\",\"action_list\",\"company_id\",\"is_admin\",\"file\",\"` + "`sensitive`" + `\",\"beizhu\",\"menu\",\"used_num\",\"sort\",\"updated_at\"]`,
	"product_category_list": `[\"id\",\"label\",\"company_id\",\"add_user\",\"add_time\",\"num\",\"sort\",\"level\",\"beizhu\",\"add_uid\",\"ck_num\",\"sn_prefix\",\"sys_cate\",\"gg_prefix\",\"pid\",\"updated_at\"]`,
	"sc_jihua_items":        `[\"id\",\"add_user\",\"add_uid\",\"add_time\",\"company_id\",\"dan_id\",\"p_id\",\"` + "`require`" + `\",\"beizhu\",\"p_num\",\"jh_date\",\"item_sn\",\"shenhe\",\"disabled\",\"do_status\",\"zc_pici_num\",\"zhuan_num\",\"gx_zhuan_num\",\"zy_uid\",\"bom_ver_id\",\"ban_num\",\"jhType\",\"out_bom_ver_id\",\"out_pid\",\"sh_user\",\"sh_time\",\"shebei\",\"level_label\",\"in_pici\",\"in_pid\",\"bao_num\",\"chengBaoNum\",\"cb_id\",\"bao_num2\",\"up_user\",\"up_time\",\"op_id\",\"order_id\",\"opi_sn\",\"c_name\",\"clListIds\",\"gy_id\",\"chengxing\",\"ban_bie\",\"ryllId\",\"bao_num_raw\",\"tuhao\",\"jh_sc_pici_sn\",\"jh_sc_pici_id\",\"reya_pai_num\",\"ban_bao_num\",\"ban_bao_count\",\"updated_at\"]`,
}

type dbConfig struct {
	UserName string
	Password string
	Addr     string
	DbName   string
}

var readerDbConfig = dbConfig{
	UserName: "root",
	Password: "root",
	//Addr:     "zz_mes_mysql8:3306",
	Addr:   "host.docker.internal:15100",
	DbName: "zhu_zhou_mes",
}

var writerDbConfig = dbConfig{
	UserName: "root",
	Password: "root",
	//Addr:     "zz_mes_mysql8:3306",
	Addr:   "host.docker.internal:15100",
	DbName: "sync_zz_mes",
}

// INSERT INTO `datax_web_db`.`job_project` (`id`, `name`, `description`, `user_id`, `flag`, `create_time`, `update_time`) VALUES (1, 'mysql数据同步', 'mysql数据同步', 1, 1, '2026-05-06 16:57:03', '2026-05-06 16:57:03');

// TRUNCATE TABLE `job_info`;

// TRUNCATE TABLE `job_log`;

// update job_info set trigger_status = 1 where project_id = 1;

func main() {
	file, err := os.OpenFile(fmt.Sprintf("./datax_sql_%d.sql", time.Now().Unix()), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("文件打开失败：%v\n", err)
		return
	}

	defer file.Close()

	firstSql := "INSERT INTO `job_info` (`job_group`, `job_cron`, `job_desc`, `project_id`, `add_time`, `update_time`, `user_id`, `alarm_email`, `executor_route_strategy`, `executor_handler`, `executor_param`, `executor_block_strategy`, `executor_timeout`, `executor_fail_retry_count`, `glue_type`, `glue_source`, `glue_remark`, `glue_updatetime`, `child_jobid`, `trigger_status`, `trigger_last_time`, `trigger_next_time`, `job_json`, `replace_param`, `jvm_param`, `inc_start_time`, `partition_info`, `last_handle_code`, `replace_param_type`, `reader_table`, `primary_key`, `inc_start_id`, `increment_type`, `datasource_id`) VALUES "

	_, err = file.WriteString(firstSql + "\n")
	if err != nil {
		fmt.Printf("写入文件失败：%v\n", err)
		return
	}

	nowTime := time.Now().Format(time.DateTime)

	for k, table := range tableList {
		totalSlots := 5 * 60 // 5分钟窗口 = 300个(分,秒)槽位
		slot := k * totalSlots / len(tableList)
		minuteOffset := slot / 60 // 0~4
		second := slot % 60       // 0~59
		// 每5分钟执行一次，但起始分/秒按表均分打散
		cronExpr := fmt.Sprintf("%02d %d/5 * * * ? *", second, minuteOffset)

		// 替换查询字段
		readerColumn := defaultReaderColumn
		if _, ok := readerTableColumn[table]; ok {
			readerColumn = readerTableColumn[table]
		}

		// 替换写入字段
		writeColumn := defaultWriteColumn
		if _, ok := writeTableColumn[table]; ok {
			writeColumn = writeTableColumn[table]
		}

		// 增量查询sql
		querySql := fmt.Sprintf(`select %s from %s where updated_at >= date_sub(curdate(),interval 1 day)`, readerColumn, table)

		// datax运行脚本
		jobJson := fmt.Sprintf(`{\"job\":{\"setting\":{\"speed\":{\"channel\":1}},\"content\":[{\"reader\":{\"name\":\"mysqlreader\",\"parameter\":{\"username\":\"%s\",\"password\":\"%s\",\"splitPk\":\"id\",\"connection\":[{\"jdbcUrl\":[\"jdbc:mysql://%s/%s?useSSL=false&allowPublicKeyRetrieval=true&characterEncoding=utf8\"],\"querySql\":[\"%s\"]}]}},\"writer\":{\"name\":\"mysqlwriter\",\"parameter\":{\"writeMode\":\"update\",\"username\":\"%s\",\"password\":\"%s\",\"connection\":[{\"jdbcUrl\":\"jdbc:mysql://%s/%s?useSSL=false&allowPublicKeyRetrieval=true&characterEncoding=utf8\",\"table\":[\"%s\"]}],\"column\":%s}}}]}}`,
			readerDbConfig.UserName, readerDbConfig.Password, readerDbConfig.Addr, readerDbConfig.DbName,
			querySql,
			writerDbConfig.UserName, writerDbConfig.Password, writerDbConfig.Addr, writerDbConfig.DbName,
			table, writeColumn)

		sql := fmt.Sprintf(`(1, '%s', '%s', 1, '%s', '%s', 1, '', 'ROUND', 'executorJobHandler', '', 'COVER_EARLY', 4, 0, 'BEAN', NULL, NULL, '%s', '', 0, 0, 0, '%s', '', '', NULL, '', 200, 'Timestamp', '', '', NULL, 0, 0)`,
			cronExpr, table, nowTime, nowTime, nowTime, jobJson)

		if k != len(tableList)-1 {
			sql += ","
		} else {
			sql += ";"
		}
		_, err = file.WriteString(sql + "\n")
		if err != nil {
			fmt.Printf("写入文件失败：%v\n", err)
			return
		}
	}

	fmt.Println("finished")
}
