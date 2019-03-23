#!/usr/bin/env bash

rpm_list_names() {
	rpm -qa --qf="%{NAME}\n"
}

rpm_query_pkg_xml() {
	rpm -q --xml ${1}
}

rpm_query_xml() {
	echo "<rpmHeaders>"
	rpm_list_names | while read line ; do
		rpm_query_pkg_xml "${line}"
	done
	echo "</rpmHeaders>"
}

# xmlstarlet can be your friend here
# see http://xmlstar.sourceforge.net/doc/xmlstarlet.pdf for examples
rpm_query_xml
