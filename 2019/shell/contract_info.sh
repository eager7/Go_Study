#!/bin/sh
if [ $# -lt 2 ]; then
	echo "USAGE:`basename $0` [account] [audit id]"
	exit 1
fi

sql="replace into eos_park.t_contract_info(\`account\`,\`consistency\`,\`audit_state\`) values('`basename $1`',2,1);replace into eos_park.t_contract_audit_info(\`task_state\`,\`account_name\`,\`auditor_id\`) values(1,'`basename $1`',`basename $2`);commit;"

echo ${sql}

mysql -uroot -h172.31.97.4 -p -e "${sql}"