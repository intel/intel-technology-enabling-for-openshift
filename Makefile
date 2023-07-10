terrascan:
	@ls */*.yaml | while read f ; \
	do \
		echo "==== $$f ====" ; \
		terrascan scan -v --show-passed -d $$(dirname $$f) -i k8s --severity high \
		--skip-rules 'AC_K8S_0051,AC_K8S_0076,AC_K8S_0087' \
		|| exit $$? ; \
	done > iocp-terrascan-result.txt

clean-terrascan:
	rm -rf iocp-terrascan-result.txt