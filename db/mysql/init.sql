
-- 系统字典参数
INSERT INTO `sys_dic` (`id`, `create_time`, `update_time`, `parent_id`, `name`, `value`, `status`, `remark`) VALUES (1, '2019-12-9 00:08:10', '2019-12-9 00:08:10', 0, 'LOG_KEEP_SAVE_DAY_KEY', '30', 1, '默认日志保存天数');


-- 初始化部门
INSERT INTO `sys_department` (`id`, `create_time`, `update_time`, `name`, `parent_id`, `status`, `order_num`, `remark`) VALUES (1, '2019-12-5 09:54:32', '2019-12-5 09:54:37', '部门1', 0, 1, 0, NULL);
INSERT INTO `sys_department` (`id`, `create_time`, `update_time`, `name`, `parent_id`, `status`, `order_num`, `remark`) VALUES (2, '2019-12-5 09:54:35', '2019-12-5 09:54:41', '部门1-1', 1, 1, 0, NULL);

-- 初始化用户
INSERT INTO `sys_user` (`id`, `create_time`, `update_time`, `name`, `username`, `password`, `phone`, `email`, `head_img`, `department_id`, `status`, `remark`, `nick_name`) VALUES (1, '2019-12-5 09:54:16', '2019-12-5 09:54:21', 'admin', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '110', '11@qq.com', NULL, 1, 1, '超级管理员', '超级管理员');
INSERT INTO `sys_user` (`id`, `create_time`, `update_time`, `name`, `username`, `password`, `phone`, `email`, `head_img`, `department_id`, `status`, `remark`, `nick_name`) VALUES (2, '2019-12-5 09:54:19', '2019-12-5 09:54:24', 'test', 'test', 'e10adc3949ba59abbe56e057f20f883e', '120', '120@qq.com', NULL, 2, 1, '测试员', 'test01');


-- 初始化菜单

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (1, '2019-09-11 11:14:44.000000', '2019-11-18 15:56:36.000000', NULL, '工作台', '/', NULL, 0, 'icon-gongzuotai', 1, NULL, 1);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (2, '2019-09-11 11:14:47.000000', '2019-09-18 15:37:18.000000', NULL, '系统', '/sys', NULL, 0, 'icon-systemfill', 2, NULL, 1);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (3, '2019-09-11 23:19:57.053631', '2019-09-12 15:53:39.000000', 18, '菜单列表', '/sys/menu', NULL, 1, 'icon-menu', 2, 'views/system/menu/index.vue', 1);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (4, '2019-09-12 00:19:27.325922', '2019-09-12 00:19:27.325922', 3, '新增', NULL, 'sys:menu:add', 2, NULL, 1, NULL, 0);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (5, '2019-09-12 00:19:51.101514', '2019-09-12 00:19:51.101514', 3, '删除', NULL, 'sys:menu:delete', 2, NULL, 2, NULL, 0);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (6, '2019-09-12 00:20:05.150021', '2019-09-12 00:20:05.150021', 3, '修改', NULL, 'sys:menu:update', 2, NULL, 3, NULL, 0);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (7, '2019-09-12 00:20:19.341206', '2019-09-12 00:20:19.341206', 3, '查询', NULL, 'sys:menu:page,sys:menu:list,sys:menu:info', 2, NULL, 4, NULL, 0);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (8, '2019-09-12 00:31:25.334489', '2019-09-15 23:45:57.000000', 18, '用户列表', '/sys/user', NULL, 1, 'icon-user', 1, 'views/system/user/index.vue', 1);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (9, '2019-09-12 00:32:21.000000', '2019-11-09 22:22:19.000000', 8, '新增', NULL, 'sys:user:add', 2, NULL, 1, NULL, 0);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (10, '2019-09-12 00:32:34.578645', '2019-09-12 00:32:34.578645', 8, '删除', NULL, 'sys:user:delete', 2, NULL, 2, NULL, 0);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (11, '2019-09-12 00:32:47.291027', '2019-09-12 00:32:47.291027', 8, '修改', NULL, 'sys:user:update', 2, NULL, 3, NULL, 0);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (12, '2019-09-12 00:33:00.053653', '2019-09-12 00:33:00.053653', 8, '查询', NULL, 'sys:user:page,sys:user:list,sys:user:info', 2, NULL, 4, NULL, 0);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (13, '2019-09-12 00:34:01.141743', '2019-09-15 23:47:27.000000', 18, '角色列表', '/sys/role', NULL, 1, 'icon-role', 3, 'views/system/role/index.vue', 1);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (14, '2019-09-12 00:34:23.459460', '2019-09-12 00:34:23.459460', 13, '新增', NULL, 'sys:role:add', 2, NULL, 1, NULL, 0);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (15, '2019-09-12 00:34:40.523471', '2019-09-12 00:34:40.523471', 13, '删除', NULL, 'sys:role:delete', 2, NULL, 2, NULL, 0);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (16, '2019-09-12 00:34:53.306692', '2019-09-12 00:34:53.306692', 13, '修改', NULL, 'sys:role:update', 2, NULL, 3, NULL, 0);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (17, '2019-09-12 00:35:05.024307', '2019-09-12 00:35:05.024307', 13, '查询', NULL, 'sys:role:page,sys:role:list,sys:role:info', 2, NULL, 4, NULL, 0);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (18, '2019-09-12 15:52:44.342387', '2019-09-15 22:11:56.000000', 2, '权限管理', NULL, NULL, 0, 'icon-permission', 1, NULL, 0);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (19, '2019-09-12 17:19:16.827254', '2019-09-15 22:12:08.000000', 2, '系统监控', NULL, NULL, 0, 'icon-monitor', 2, NULL, 0);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (20, '2019-09-12 17:35:51.000000', '2019-11-26 23:46:53.000000', 19, '请求日志', '/sys/log', '', 1, 'icon-log', 1, 'views/system/log/index.vue', 0);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (21, '2019-09-12 17:37:03.000000', '2019-11-18 14:20:27.000000', 20, '权限', NULL, 'sys:log:page,sys:log:clear,sys:log:get-keep,sys:log:set-keep', 2, NULL, 1, NULL, 0);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (22, '2019-11-07 14:22:34.000000', '2019-11-09 22:11:18.000000', 23, 'crud 示例', '/crud', NULL, 1, 'icon-radioboxfill', 1, 'views/test/crud.vue', 1);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (23, '2019-11-07 22:36:57.000000', '2019-11-11 15:21:10.000000', 1, '组件库', '/ui-lib', NULL, 0, 'icon-activityfill', 2, NULL, 1);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (24, '2019-11-08 09:35:08.000000', '2019-11-08 19:44:44.000000', NULL, '教程', '/tutorial', NULL, 0, 'icon-weibiaoti4', 3, NULL, 1);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (25, '2019-11-08 09:35:53.000000', '2019-11-08 09:37:02.000000', 24, '文档', '/tutorial/doc', NULL, 1, 'icon-favorfill', 0, 'https://docs.cool-admin.com/#/', 1);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (26, '2019-11-09 22:11:13.000000', '2019-11-09 22:11:34.000000', 23, 'Quill富文本', '/editor-quill', NULL, 1, 'icon-radioboxfill', 3, 'views/test/editor-quill.vue', 1);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (27, '2019-11-11 15:21:00.457911', '2019-11-11 15:21:00.457911', 1, '组件预览', '/ui-components', NULL, 1, 'icon-activityfill', 1, 'views/workbench/ui-components', 1);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (28, '2019-11-18 16:50:27.676184', '2019-11-18 16:50:27.676184', 8, '部门列表', NULL, 'sys:department:list', 2, NULL, 0, NULL, 1);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (29, '2019-11-18 16:50:45.783945', '2019-11-18 16:50:45.783945', 8, '新增部门', NULL, 'sys:department:add', 2, NULL, 0, NULL, 1);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (30, '2019-11-18 16:50:59.241883', '2019-11-18 16:50:59.241883', 8, '更新部门', NULL, 'sys:department:update', 2, NULL, 0, NULL, 1);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (31, '2019-11-18 16:51:13.304712', '2019-11-18 16:51:13.304712', 8, '删除部门', NULL, 'sys:department:delete', 2, NULL, 0, NULL, 1);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (32, '2019-11-18 17:49:35.009792', '2019-11-18 17:49:35.009792', 8, '部门排序', NULL, 'sys:department:order', 2, NULL, 0, NULL, 1);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (33, '2019-11-18 23:59:21.775467', '2019-11-18 23:59:21.775467', 8, '用户转移', NULL, 'sys:user:move', 2, NULL, 0, NULL, 1);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (34, '2019-11-26 23:48:16.000000', '2019-11-26 23:48:29.000000', 19, '服务状态', '/sys/perf', NULL, 1, 'icon-monitor', 2, 'views/system/perf/index.vue', 0);
INSERT INTO `sys_menu`(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`) VALUES (35, '2019-11-26 23:50:13.059599', '2019-11-26 23:50:13.059599', 34, '记录', NULL, 'sys:info:record', 2, NULL, 0, NULL, 1);


-- 新增 菜单 用户转移部门
INSERT INTO `sys_menu`
(`id`, `create_time`, `update_time`, `parent_id`, `name`, `url`, `perms`,`type`, `icon`, `order_num`, `view_path`, `keep_alive`)
VALUES
 (36, '2019-11-26 23:50:13.059599', '2019-11-26 23:50:13.059599', 8, '用户转移', NULL, 'sys:user:move', 2, NULL, 0, NULL, 1);


-- 初始化角色
INSERT INTO `sys_role` (`id`, `create_time`, `update_time`, `name`, `status`, `remark`) VALUES (1, '2019-12-5 09:30:30', '2019-12-5 09:30:38', '超级管理员', 1, '超级管理员');


-- 初始化用户跟角色关系
INSERT INTO `sys_user_role` (`id`, `create_time`, `update_time`, `user_id`, `role_id`, `remark`) VALUES (1, '2019-12-5 09:57:25', '2019-12-5 09:57:27', 1, 1, NULL);


-- 初始化角色对应权限关系
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 1);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 2);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 3);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 4);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 5);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 6);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 7);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 8);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 9);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 10);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 11);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 12);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 13);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 14);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 15);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 16);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 17);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 18);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 19);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 20);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 21);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 22);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 23);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 24);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 25);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 26);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 27);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 28);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 29);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 30);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 31);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 32);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 33);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 34);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 35);
INSERT INTO `sys_role_menu` ( `role_id`, `menu_id`) VALUES (1, 36);



